package account

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetTokenDetailListReq struct {
		g.Meta `path:"/v1/account/token/list" tags:"User" method:"get" summary:"User token list"`
		model.AuthorizationRequired
		Locale string
	}
	GetTokenDetailListRes struct {
		List []*model.RefreshTokenDetail `json:"list"`
	}
)
