package token

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetUserInfoReq struct {
		g.Meta `path:"/user/userinfo" tags:"SDK" method:"get" summary:"Get User Info"`
		model.AppAuthorizationRequired
	}
	GetUserInfoRes model.UserAccount
)
