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
	GetInfoRes model.UserAccount
)
