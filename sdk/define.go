package sdk

import (
	"time"
)

type (
	TokenInfo struct {
		AccountId       string    `json:"account_id"`
		SpaceId         int64     `json:"space_id"`
		UserId          string    `json:"user_id"`
		ExpireAt        time.Time `json:"expire_at"`
		RefreshTokenKey string    `json:"refresh_token_key"`
		AccessToken     string    `json:"access_token"`
	}
)
