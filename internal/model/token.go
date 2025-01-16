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
	RefreshTokenClaims struct {
		jwt.RegisteredClaims
		RefreshTokenExtraData
	}
	RefreshTokenExtraData struct {
		ClientIP string `json:"cip,omitempty"`
		From     string `json:"frm,omitempty"`
		UA       string `json:"uag,omitempty"`
	}
	TokenInfo struct {
		AccountId       string    `json:"account_id"`
		SpaceId         int64     `json:"space_id"`
		UserId          string    `json:"user_id"`
		ExpireAt        time.Time `json:"expire_at"`
		RefreshTokenKey string    `json:"refresh_token_key"`
		AccessToken     string    `json:"access_token"`
	}
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
