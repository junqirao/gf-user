package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-user/internal/consts"
	"gf-user/internal/model/do"
)

type (
	AccountRegisterInput struct {
		Account       string         `json:"account"`
		Password      string         `json:"password"`
		Type          int            `json:"type"`
		Status        int            `json:"status"`
		Name          string         `json:"name"`
		Email         string         `json:"email"`
		Avatar        string         `json:"avatar"`
		Extra         map[string]any `json:"extra"`
		Administrator bool           `json:"-"`
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
		Spaces    []int64     `json:"spaces"`
	}

	AccountLoginInput struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Nonce    string `json:"nonce"`
		From     string `json:"from"`
	}

	UserAccountLoginInfo struct {
		*UserAccount
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	AccountModifyPasswordInput struct {
		MFACodeRequired
		Old   string `json:"old"`
		New   string `json:"new"`
		Nonce string `json:"nonce"`
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
		HasMFA:    len(account.Mfa) > 0,
	}
	if len(sp) > 0 {
		for i, space := range sp {
			if gconv.Int64(space.Id) == gconv.Int64(user.Space) {
				ua.SpaceInfo = sp[i]
			}
			ua.Spaces = append(ua.Spaces, space.Id)
		}
	}
	extra := gconv.Map(account.Extra)
	delete(extra, consts.AccountExtraKeyAdminCode)
	ua.Extra = extra
	return ua
}
