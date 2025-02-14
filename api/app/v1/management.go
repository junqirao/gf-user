package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	CreateAppReq struct {
		g.Meta `path:"/v1/app" tags:"App" method:"put" summary:"Create App"`
		model.AuthorizationRequired
		model.CreateAppInput
	}
	CreateAppRes model.CreateAppOutput

	UpdateAppReq struct {
		g.Meta `path:"/v1/app/{app_id}" tags:"App" method:"post" summary:"Update App"`
		model.AuthorizationRequired
		model.UpdateAppInput
	}
	UpdateAppRes struct{}

	RemoveAppReq struct {
		g.Meta `path:"/v1/app/{app_id}" tags:"App" method:"delete" summary:"Remove App"`
		model.AuthorizationRequired
		AppId string `json:"app_id" in:"path"`
	}
	RemoveAppRes struct{}

	ListAppReq struct {
		g.Meta `path:"/v1/app" tags:"App" method:"get" summary:"List App"`
		model.AuthorizationRequired
	}
	ListAppRes struct {
		List []*model.AppInfo `json:"list"`
	}

	GetAppOneReq struct {
		g.Meta `path:"/v1/app/{app_id}" tags:"App" method:"get" summary:"Get App One"`
		model.AuthorizationRequired
		AppId string `json:"app_id" in:"path"`
	}
	GetAppOneRes model.AppInfo
)
