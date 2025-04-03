package token

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
	"github.com/junqirao/gocomponents/kvdb"

	"gf-user/internal/consts"
	"gf-user/internal/logic/config"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func init() {
	ctx := gctx.GetInitCtx()
	config.MustInit(ctx)
	service.RegisterToken(newSToken(ctx))
}

type (
	sToken struct {
	}
	refreshToken struct {
		Key      string
		ExpireAt int64
	}
)

func newSToken(_ context.Context) *sToken {
	return &sToken{}
}

func (t sToken) GenerateAccessToken(ctx context.Context, user *model.UserAccount, extra model.RefreshTokenExtraData) (accessToken string, refreshToken string, err error) {
	mu, err := kvdb.NewMutex(ctx, fmt.Sprintf("user_token_handler_%v", user.Id))
	if err != nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	cfg := service.Config().GetTokenConfig(ctx)

	refreshTokenKey := grand.S(8)
	if accessToken, err = t.signAccessToken(cfg, user, refreshTokenKey); err != nil {
		return
	}

	var exp time.Time
	refreshToken, exp, err = t.signRefreshToken(cfg, user, refreshTokenKey, extra)
	if err != nil {
		return
	}

	key := t.getUserRefreshTokenKey(user.Id)
	if _, err = g.Redis().ZAdd(ctx, key, &gredis.ZAddOption{}, gredis.ZAddMember{Score: float64(exp.Unix()), Member: refreshTokenKey}); err != nil {
		return
	}

	if err = t.storeExtraInfo(ctx, cfg, user.Id, refreshTokenKey, extra); err != nil {
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
	cfg := service.Config().GetTokenConfig(ctx)
	token, err := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})).
		ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.TokenKey), nil
		})
	err = gerror.Cause(err)
	if err != nil {
		return
	}
	if !token.Valid || len(claims.Audience) == 0 {
		err = code.ErrInvalidToken
		return
	}
	// validate appid if exists
	if claims.AppId != "" {
		var app *model.AppInfo
		app, err = service.App().Info(ctx, claims.AppId)
		if err != nil {
			return
		}
		// check space id
		if app.Space != gconv.Int(claims.SpaceId) {
			err = code.ErrAppSpaceNotAllowed
			return
		}
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
				AppId:           claims.AppId,
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
			ext := model.RefreshTokenExtraData{
				From: claims.From,
			}
			if req := ghttp.RequestFromCtx(ctx); req != nil {
				ext.ClientIP = req.GetClientIp()
				ext.UA = req.UserAgent()
			}
			newAccessToken, newRefreshToken, err = t.GenerateAccessToken(ctx, user, ext)
			return
		}
	}

	err = code.ErrRefreshTokenNotFound
	return
}

func (t sToken) GenerateAppToken(ctx context.Context, appId string, user *model.UserAccount, claims *model.RefreshTokenClaims) (accessToken string, err error) {
	mu, err := kvdb.NewMutex(ctx, fmt.Sprintf("user_token_handler_%v", user.Id))
	if err != nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	cfg := service.Config().GetTokenConfig(ctx)

	key := t.getUserRefreshTokenKey(user.Id)
	rts, err := t.getUserRefreshTokens(ctx, key)
	if err != nil {
		return
	}
	for _, rt := range rts {
		if rt.Key == claims.Subject {
			if accessToken, err = t.signAccessToken(cfg, user, rt.Key, appId); err != nil {
				return
			}
			return
		}
	}

	err = code.ErrRefreshTokenNotFound
	return
}

func (t sToken) RemoveRefreshToken(ctx context.Context, accountId string, claims *model.RefreshTokenClaims) (err error) {
	err = t.removeRefreshToken(ctx, accountId, claims.Subject)
	return
}

func (t sToken) ParseRefreshToken(ctx context.Context, refreshToken string) (claims *model.RefreshTokenClaims, err error) {
	cfg := service.Config().GetTokenConfig(ctx)
	claims = new(model.RefreshTokenClaims)
	token, err := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})).
		ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.TokenKey), nil
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
	// remove expired
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
	cfg := service.Config().GetTokenConfig(ctx)
	// remove over limits
	if cfg.RefreshTokenLimit <= 0 {
		return
	}
	overCount := int64(len(rts)) - expCount - cfg.RefreshTokenLimit
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

func (t sToken) signAccessToken(cfg *model.UserTokenConfig, user *model.UserAccount, refreshTokenKey string, appId ...string) (accessToken string, err error) {
	ts := time.Now()
	var app string
	if len(appId) > 0 {
		app = appId[0]
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256,
		&model.AccessTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  []string{gconv.String(user.Id)},
				ExpiresAt: jwt.NewNumericDate(ts.Add(time.Second * time.Duration(cfg.AccessTokenExpire))), // 过期时间
				IssuedAt:  jwt.NewNumericDate(ts),
				Issuer:    consts.DefaultTokenIssuer, // 签发人
				NotBefore: jwt.NewNumericDate(ts),
				Subject:   refreshTokenKey,
			},
			SpaceId: gconv.String(user.SpaceInfo.Id),
			UserId:  gconv.String(user.UserInfo.Id),
			AppId:   app,
		},
	).SignedString([]byte(cfg.TokenKey))
}

func (t sToken) signRefreshToken(cfg *model.UserTokenConfig, user *model.UserAccount, refreshTokenKey string, extra model.RefreshTokenExtraData) (refreshToken string, exp time.Time, err error) {
	ts := time.Now()
	exp = ts.Add(time.Second * time.Duration(cfg.RefreshTokenExpire))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
		&model.RefreshTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  []string{gconv.String(user.Id)},
				ExpiresAt: jwt.NewNumericDate(exp), // 过期时间
				IssuedAt:  jwt.NewNumericDate(ts),
				Issuer:    consts.DefaultTokenIssuer, // 签发人
				NotBefore: jwt.NewNumericDate(ts),
				Subject:   refreshTokenKey,
			},
			RefreshTokenExtraData: extra,
		},
	).SignedString([]byte(cfg.TokenKey))
	return
}

func (t sToken) getUserRefreshTokenKey(id any) string {
	return fmt.Sprintf("user:refresh_token:%v", id)
}

func (t sToken) getUserRefreshTokenExtraDataKey(id any, key string) string {
	return fmt.Sprintf("user:refresh_token:extra:%v:%s", id, key)
}

func (t sToken) storeExtraInfo(ctx context.Context,
	cfg *model.UserTokenConfig,
	accountId any, key string,
	extra model.RefreshTokenExtraData) (err error) {
	ext := cfg.RefreshTokenExpire
	k := t.getUserRefreshTokenExtraDataKey(accountId, key)
	data := gbase64.EncodeString(gconv.String(extra))
	_, err = g.Redis().Set(ctx, k, data, gredis.SetOption{TTLOption: gredis.TTLOption{EX: &ext}})
	return
}

func (t sToken) removeRefreshToken(ctx context.Context, accountId, key string) (err error) {
	_, err = g.Redis().ZRem(ctx, t.getUserRefreshTokenKey(accountId), key)
	if err != nil {
		return
	}
	_, err = g.Redis().Unlink(ctx, t.getUserRefreshTokenExtraDataKey(accountId, key))
	return
}
