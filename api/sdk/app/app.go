package app

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetAppInfoReq struct {
		g.Meta `path:"/app/info" tags:"SDK" method:"get" summary:"Get app info"`
		model.AppAuthorizationRequired
	}
	GetAppInfoRes model.AppInfo
)
