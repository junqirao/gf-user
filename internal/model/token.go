package model

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthorizationRequired struct {
		Authorization string `p:"Authorization" in:"header" v:"required" `
	}
	AccessTokenClaims struct {
		jwt.RegisteredClaims
		SpaceId string `json:"sid,omitempty"`
		UserId  string `json:"uid,omitempty"`
	}
	RefreshTokenClaims = jwt.RegisteredClaims
	UserTokenConfig    struct {
		AccessTokenExpire  int64  `json:"access_token_expire" default:"7200"`
		RefreshTokenExpire int64  `json:"refresh_token_expire" default:"2592000"`
		RefreshTokenLimit  int64  `json:"refresh_token_limit"`
		TokenKey           string `json:"token_key,omitempty" default:"dEfaUlTuSerT0k3nK3y"`
	}
	TokenInfo struct {
		AccountId       string    `json:"account_id"`
		SpaceId         int64     `json:"space_id"`
		UserId          string    `json:"user_id"`
		ExpireAt        time.Time `json:"expire_at"`
		RefreshTokenKey string    `json:"refresh_token_key"`
		AccessToken     string    `json:"access_token"`
	}
)

func (u UserTokenConfig) Print(ctx context.Context) {
	g.Log().Infof(ctx, "user token config: access_token_expire: %d, refresh_token_expire: %d, refresh_token_limit: %d", u.AccessTokenExpire, u.RefreshTokenExpire, u.RefreshTokenLimit)
}
