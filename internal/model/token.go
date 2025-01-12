package model

import (
	"time"

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
	TokenInfo          struct {
		AccountId       string    `json:"account_id"`
		SpaceId         int64     `json:"space_id"`
		UserId          string    `json:"user_id"`
		ExpireAt        time.Time `json:"expire_at"`
		RefreshTokenKey string    `json:"refresh_token_key"`
		AccessToken     string    `json:"access_token"`
	}
)
