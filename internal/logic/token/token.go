package token

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
	"github.com/junqirao/gocomponents/kvdb"

	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func init() {
	service.RegisterToken(newSToken(gctx.GetInitCtx()))
}

type (
	sToken struct {
		mu  sync.Locker
		cfg model.UserTokenConfig
	}
	refreshToken struct {
		Key      string
		ExpireAt int64
	}
)

func newSToken(ctx context.Context) *sToken {
	cfg := model.UserTokenConfig{}
	v, err := g.Cfg().Get(ctx, "token")
	if err == nil {
		_ = v.Struct(&cfg)
	}
	if cfg.AccessTokenExpire <= 0 {
		cfg.AccessTokenExpire = consts.DefaultAccessTokenExpire
	}
	if cfg.RefreshTokenExpire <= 0 {
		cfg.RefreshTokenExpire = consts.DefaultRefreshTokenExpire
	}
	if cfg.TokenKey == "" {
		cfg.TokenKey = consts.DefaultTokenKey
	}
	cfg.Print(ctx)
	mutex, err := kvdb.NewMutex(ctx, "user_token_handler")
	if err != nil {
		return nil
	}
	return &sToken{mu: mutex, cfg: cfg}
}

func (t sToken) GenerateAccessToken(ctx context.Context, user *model.UserAccount) (accessToken string, refreshToken string, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	refreshTokenKey := grand.S(8)
	if accessToken, err = t.signAccessToken(user, refreshTokenKey); err != nil {
		return
	}
	var exp time.Time
	refreshToken, exp, err = t.signRefreshToken(user, refreshTokenKey)
	if err != nil {
		return
	}

	key := t.getUserRefreshTokenKey(user.Id)
	if _, err = g.Redis().ZAdd(ctx, key, &gredis.ZAddOption{}, gredis.ZAddMember{Score: float64(exp.Unix()), Member: refreshTokenKey}); err != nil {
		return
	}

	rts, err := t.getUserRefreshTokens(ctx, key)
	if err != nil {
		return
	}
	err = t.removeInvalidRefreshTokens(ctx, key, rts)
	return
}

func (t sToken) ValidAccessToken(ctx context.Context, accessToken string) (tokenInfo *model.TokenInfo, err error) {
	claims := new(model.AccessTokenClaims)
	defer func() {
		if err != nil && !errors.Is(err, code.ErrInvalidToken) {
			err = code.ErrInvalidToken.WithDetail(err.Error())
		}
	}()
	token, err := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})).
		ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(t.cfg.TokenKey), nil
		})
	err = gerror.Cause(err)
	if err != nil {
		return
	}
	if !token.Valid || len(claims.Audience) == 0 {
		err = code.ErrInvalidToken
		return
	}
	tokens, err := t.getUserRefreshTokens(ctx, t.getUserRefreshTokenKey(claims.Audience[0]))
	if err != nil {
		return
	}
	for _, v := range tokens {
		if v.Key == claims.Subject {
			tokenInfo = &model.TokenInfo{
				AccountId:       claims.Audience[0],
				SpaceId:         gconv.Int64(claims.SpaceId),
				UserId:          claims.UserId,
				ExpireAt:        claims.ExpiresAt.Time,
				RefreshTokenKey: claims.Subject,
				AccessToken:     accessToken,
			}
			return
		}
	}
	err = errors.New("invalid refresh token")
	return
}

// GetTokenInfoFromCtx only work under middleware.AuthToken
func (t sToken) GetTokenInfoFromCtx(ctx context.Context) (tokenInfo model.TokenInfo) {
	v := ctx.Value(consts.CtxKeyTokenInfo)
	if v == nil {
		return
	}
	ti, ok := v.(*model.TokenInfo)
	if ok {
		tokenInfo = *ti
	}
	return
}

func (t sToken) RefreshToken(ctx context.Context, user *model.UserAccount, claims *model.RefreshTokenClaims) (newAccessToken, newRefreshToken string, err error) {
	key := t.getUserRefreshTokenKey(user.Id)
	rts, err := t.getUserRefreshTokens(ctx, key)
	if err != nil {
		return
	}

	for _, rt := range rts {
		if rt.Key == claims.Subject {
			if _, err = g.Redis().ZRem(ctx, key, rt.Key); err != nil {
				return
			}
			newAccessToken, newRefreshToken, err = t.GenerateAccessToken(ctx, user)
			return
		}
	}

	err = code.ErrInvalidToken.WithDetail("refresh token not found")
	return
}

func (t sToken) ParseRefreshToken(_ context.Context, refreshToken string) (claims *model.RefreshTokenClaims, err error) {
	claims = new(model.RefreshTokenClaims)
	token, err := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})).
		ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(t.cfg.TokenKey), nil
		})
	if err != nil {
		return
	}
	if !token.Valid {
		err = code.ErrInvalidToken
		return
	}
	return
}

func (t sToken) removeInvalidRefreshTokens(ctx context.Context, key string, rts []*refreshToken) (err error) {
	// 移除超时
	now := time.Now().Unix()
	var expCount int64 = 0
	for _, rt := range rts {
		if now > rt.ExpireAt {
			expCount++
			continue
		}
		break
	}
	if expCount > 0 {
		g.Log().Infof(ctx, "remove refesh token %s, expired count: %d", key, expCount)
		if _, err = g.Redis().ZRemRangeByRank(ctx, key, 0, expCount); err != nil {
			return
		}
	}
	// 移除超限
	if t.cfg.RefreshTokenLimit <= 0 {
		return
	}
	overCount := int64(len(rts)) - expCount - t.cfg.RefreshTokenLimit
	if overCount <= 0 {
		return
	}
	g.Log().Infof(ctx, "remove refesh token %s, over count: %d", key, overCount)
	// remove [0,n]
	_, err = g.Redis().ZRemRangeByRank(ctx, key, 0, overCount-1)
	return
}

func (t sToken) getUserRefreshTokens(ctx context.Context, key string) (rts []*refreshToken, err error) {
	list, err := g.Redis().ZRange(ctx, key, 0, -1, gredis.ZRangeOption{WithScores: true})
	if err != nil {
		return
	}
	for _, v := range list {
		slice := v.Slice()
		if len(slice) != 2 {
			continue
		}

		rts = append(rts, &refreshToken{
			Key:      gconv.String(slice[0]),
			ExpireAt: gconv.Int64(slice[1]),
		})
	}
	return
}

func (t sToken) signAccessToken(user *model.UserAccount, refreshTokenKey string) (accessToken string, err error) {
	ts := time.Now()
	return jwt.NewWithClaims(jwt.SigningMethodHS256,
		&model.AccessTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  []string{gconv.String(user.Id)},
				ExpiresAt: jwt.NewNumericDate(ts.Add(time.Second * time.Duration(t.cfg.AccessTokenExpire))), // 过期时间
				IssuedAt:  jwt.NewNumericDate(ts),
				Issuer:    consts.DefaultTokenIssuer, // 签发人
				NotBefore: jwt.NewNumericDate(ts),
				Subject:   refreshTokenKey,
			},
			SpaceId: gconv.String(user.SpaceInfo.Id),
			UserId:  gconv.String(user.UserInfo.Id),
		},
	).SignedString([]byte(t.cfg.TokenKey))
}

func (t sToken) signRefreshToken(user *model.UserAccount, refreshTokenKey string) (refreshToken string, exp time.Time, err error) {
	ts := time.Now()
	exp = ts.Add(time.Second * time.Duration(t.cfg.RefreshTokenExpire))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
		&model.RefreshTokenClaims{
			Audience:  []string{gconv.String(user.Id)},
			ExpiresAt: jwt.NewNumericDate(exp), // 过期时间
			IssuedAt:  jwt.NewNumericDate(ts),
			Issuer:    consts.DefaultTokenIssuer, // 签发人
			NotBefore: jwt.NewNumericDate(ts),
			Subject:   refreshTokenKey,
		},
	).SignedString([]byte(t.cfg.TokenKey))
	return
}

func (t sToken) getUserRefreshTokenKey(id any) string {
	return fmt.Sprintf("user:refresh_token:%v", id)
}
