package account

import (
	"github.com/gogf/gf/v2/frame/g"
)

type (
	ModifyNameReq struct {
		g.Meta `path:"/v1/account/name" tags:"User" method:"post" summary:"User modify name"`
		Name   string `json:"name"`
	}
	ModifyNameRes   struct{}
	ModifyAvatarReq struct {
		g.Meta `path:"/v1/account/avatar" tags:"User" method:"post" summary:"User modify avatar"`
		Avatar string `json:"avatar"`
	}
	ModifyAvatarRes   struct{}
	ModifyPasswordReq struct {
		g.Meta `path:"/v1/account/password" tags:"User" method:"post" summary:"User modify password"`
		Old    string `json:"old" dc:"md5(md5(raw)+nonce)"`
		New    string `json:"new" dc:"md5(raw)"`
		Nonce  string `json:"nonce"`
	}
	ModifyPasswordRes struct{}
)
