package user

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetInfoReq struct {
		g.Meta `path:"/v1/account/user/info" tags:"User" method:"get" summary:"User info"`
		model.AuthorizationRequired
	}
	GetInfoRes        model.UserAccount
	ModifyUserNameReq struct {
		g.Meta `path:"/v1/account/user/name" tags:"User" method:"post" summary:"User modify name"`
		model.AuthorizationRequired
		Name string `json:"name"`
	}
	ModifyUserNameRes struct{}
)
