package model

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/junqirao/gf-user/sdk"
)

type (
	AuthorizationRequired struct {
		Authorization string `p:"Authorization" in:"header" v:"required" `
	}
	AccessTokenClaims struct {
		jwt.RegisteredClaims
		SpaceId string `json:"sid,omitempty"`
		UserId  string `json:"uid,omitempty"`
		AppId   string `json:"app,omitempty"`
	}
	RefreshTokenClaims struct {
		jwt.RegisteredClaims
		RefreshTokenExtraData
	}
	RefreshTokenExtraData struct {
		ClientIP string `json:"cip,omitempty"`
		From     string `json:"frm,omitempty"`
		UA       string `json:"uag,omitempty"`
	}
	TokenInfo          = sdk.TokenInfo
	RefreshTokenDetail struct {
		Key        string `json:"key"`
		City       string `json:"city"`
		Country    string `json:"country"`
		CountryISO string `json:"country_iso"`
		IP         string `json:"ip"`
		UserAgent  string `json:"user_agent"`
		From       string `json:"from"`
		ExpireAt   int64  `json:"expire_at"`
	}
)
