package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

// Login
type (
	SetLoginConfigReq struct {
		g.Meta `path:"/v1/config/login" tags:"Config" method:"post" summary:"set mfa config"`
		SetConfig
	}
	SetLoginConfigRes struct{}

	GetLoginConfigReq struct {
		g.Meta `path:"/v1/config/login" tags:"Config" method:"get" summary:"get login config"`
		model.AuthorizationRequired
	}
	GetLoginConfigRes model.LoginConfig
)
