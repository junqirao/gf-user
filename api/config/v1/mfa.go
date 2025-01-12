package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

// MFA
type (
	SetMFAConfigReq struct {
		g.Meta `path:"/v1/config/mfa" tags:"Config" method:"post" summary:"set mfa config"`
		SetConfig
	}
	SetMFAConfigRes struct{}

	GetMFAConfigReq struct {
		g.Meta `path:"/v1/config/mfa" tags:"Config" method:"get" summary:"get mfa config"`
		model.AuthorizationRequired
	}
	GetMFAConfigRes model.MFAConfig
)
