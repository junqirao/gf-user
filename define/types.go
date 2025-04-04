package define

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	TokenInfo struct {
		AccountId       string    `json:"account_id"`
		SpaceId         int64     `json:"space_id"`
		UserId          string    `json:"user_id"`
		ExpireAt        time.Time `json:"expire_at"`
		RefreshTokenKey string    `json:"refresh_token_key"`
		AccessToken     string    `json:"access_token"`
		AppId           string    `json:"app_id,omitempty"`
	}

	Space struct {
		Id          int64          `json:"id"`
		Name        string         `json:"name"`
		Logo        string         `json:"logo"`
		LogoKey     string         `json:"logo_key"`
		IsOwner     bool           `json:"is_owner"`
		Description string         `json:"description"`
		Profile     map[string]any `json:"profile"`
		UpdateAt    *gtime.Time    `json:"update_at"`
		CreatedAt   *gtime.Time    `json:"created_at"`
	}

	AccountBrief struct {
		Id        interface{} `json:"id"` // account uuid
		Name      interface{} `json:"name"`
		Avatar    interface{} `json:"avatar"`
		AvatarKey interface{} `json:"avatar_key"`
	}

	Account struct {
		*AccountBrief
		Account   interface{} `json:"account"`            // unique account
		Password  interface{} `json:"password,omitempty"` // password hash
		Type      interface{} `json:"type"`               // 0: normal, 1: app
		Status    interface{} `json:"status"`             // 0: normal, 1: frozen
		Email     interface{} `json:"email"`
		CreatedAt *gtime.Time `json:"created_at"`
		UpdateAt  *gtime.Time `json:"update_at"`
		HasMFA    bool        `json:"has_mfa"`
		Extra     interface{} `json:"extra"`
	}
	UserAccount struct {
		*Account
		UserInfo  *UserInfo `json:"user_info"`
		SpaceInfo *Space    `json:"space_info"`
		Spaces    []int64   `json:"spaces"`
	}

	UserAccountLoginInfo struct {
		*UserAccount
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	UserInfo struct {
		Id     interface{} `json:"id"`      // user id
		Space  interface{} `json:"space"`   // space id
		Type   interface{} `json:"type"`    // 0: normal, 1: manager
		Name   interface{} `json:"name"`    //
		JoinAt *gtime.Time `json:"join_at"` // join time
	}
)
