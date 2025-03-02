package token

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	ValidateAppTokenReq struct {
		g.Meta `path:"/token/validate" tags:"SDK" method:"post" summary:"Validate Token"`
		model.AppAuthorizationRequired
	}

	ValidateAppTokenRes model.TokenInfo
)
