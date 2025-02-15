package token

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	ValidateTokenReq struct {
		g.Meta `path:"/token/validate" tags:"SDK" method:"post" summary:"Validate Token"`
		model.AppAuthorizationRequired
		AccessToken string `json:"access_token"`
	}

	ValidateTokenRes model.TokenInfo
)
