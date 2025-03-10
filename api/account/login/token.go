package login

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	RefreshTokenReq struct {
		g.Meta `path:"/v1/account/user/refresh/token" tags:"User" method:"get" summary:"User refresh token"`
		// model.AuthorizationRequired
		RefreshToken string `json:"refresh_token"`
		Space        int64  `json:"space"`
	}
	RefreshTokenRes     model.UserAccountLoginInfo
	GenerateAppTokenReq struct {
		g.Meta `path:"/v1/account/user/app/token" tags:"User" method:"get" summary:"User sign app token"`
		// model.AuthorizationRequired
		AppId        string `json:"X-App-Id" in:"header"`
		RefreshToken string `json:"refresh_token"`
		Space        int64  `json:"space"`
	}
	GenerateAppTokenRes model.UserAccountLoginInfo
)
