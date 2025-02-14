package public

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetAppOneReq struct {
		g.Meta `path:"/public/app/{app_id}" tags:"App" method:"get" summary:"Get App One"`
		AppId  string `json:"app_id" in:"path"`
	}
	GetAppOneRes model.AppInfo
)
