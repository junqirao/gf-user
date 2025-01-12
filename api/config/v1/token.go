package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	SetTokenConfigReq struct {
		g.Meta `path:"/v1/config/token" tags:"Config" method:"post" summary:"set token config"`
		model.AuthorizationRequired
		Content *model.UserTokenConfig `json:"content"`
	}
	SetTokenConfigRes struct{}

	GetTokenConfigReq struct {
		g.Meta `path:"/v1/config/token" tags:"Config" method:"get" summary:"get token config"`
		model.AuthorizationRequired
	}
	GetTokenConfigRes model.UserTokenConfig
)
