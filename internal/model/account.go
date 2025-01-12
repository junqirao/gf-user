package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-user/internal/model/do"
)

type (
	AccountRegisterInput struct {
		Account  string         `json:"account"`
		Password string         `json:"password"`
		Type     int            `json:"type"`
		Status   int            `json:"status"`
		Name     string         `json:"name"`
		Email    string         `json:"email"`
		Avatar   string         `json:"avatar"`
		Extra    map[string]any `json:"extra"`
	}
	UserAccount struct {
		Id        interface{} `json:"id"`                              // account uuid
		Account   interface{} `json:"account"`                         // unique account
		Password  interface{} `json:"password,omitempty"`              // password hash
		Type      interface{} `json:"type" mapping:"account_type"`     // 0: normal, 1: app
		Status    interface{} `json:"status" mapping:"account_status"` // 0: normal, 1: frozen
		Name      interface{} `json:"name"`
		Email     interface{} `json:"email"`
		Avatar    interface{} `json:"avatar"`
		AvatarKey interface{} `json:"avatar_key"`
		CreatedAt *gtime.Time `json:"created_at"`
		UpdateAt  *gtime.Time `json:"update_at"`
		UserInfo  *UserInfo   `json:"user_info"`
		SpaceInfo *Space      `json:"space_info"`
		Extra     interface{} `json:"extra"`
		HasMFA    bool        `json:"has_mfa"`
	}

	AccountLoginInput struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Nonce    string `json:"nonce"`
	}

	UserAccountLoginInfo struct {
		*UserAccount
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func NewUserAccount(account *do.Account, user *do.User, sp ...*Space) *UserAccount {
	ua := &UserAccount{
		Id:        account.Id,
		Account:   account.Account,
		Type:      account.Type,
		Status:    account.Status,
		Name:      account.Name,
		Email:     account.Email,
		Avatar:    account.Avatar,
		AvatarKey: account.Avatar,
		UserInfo: &UserInfo{
			Id:     user.Id,
			Space:  user.Space,
			Type:   user.Type,
			Name:   user.Name,
			JoinAt: user.CreatedAt,
		},
		CreatedAt: account.CreatedAt,
		UpdateAt:  account.UpdateAt,
		Extra:     gconv.Map(account.Extra),
		HasMFA:    len(account.Mfa) > 0,
	}
	if len(sp) > 0 {
		ua.SpaceInfo = sp[0]
	}
	return ua
}
