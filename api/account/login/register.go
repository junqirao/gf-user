package login

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	RegisterReq struct {
		g.Meta   `path:"/v1/account/user/register" tags:"User" method:"post" summary:"User register"`
		Account  string `json:"account"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
	}
	RegisterRes                   model.UserAccount
	RegisterSuperAdministratorReq struct {
		g.Meta   `path:"/v1/account/user/register/super-admin" tags:"User" method:"post" summary:"User register super administrator"`
		Account  string `json:"account"`
		Password string `json:"password"`
		Code     string `json:"code"`
		Nonce    string `json:"nonce"`
	}
	RegisterSuperAdministratorRes model.UserAccount
)
