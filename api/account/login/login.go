package login

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	UserLoginReq struct {
		g.Meta   `path:"/v1/account/user/login" tags:"User" method:"post" summary:"User login"`
		Account  string `json:"account"`
		Password string `json:"password"`
		Nonce    string `json:"nonce"`
		From     string `json:"from"`
	}
	UserLoginRes model.UserAccountLoginInfo

	GetLoginConfigReq struct {
		g.Meta `path:"/v1/account/login/config" tags:"User" method:"get" summary:"get login config"`
	}
	GetLoginConfigRes model.LoginConfig

	UserLogoutReq struct {
		g.Meta       `path:"/v1/account/user/logout" tags:"User" method:"post" summary:"User logout"`
		RefreshToken string `json:"refresh_token" v:"required"`
	}
	UserLogoutRes struct{}
)
