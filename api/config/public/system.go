package public

import (
	"github.com/gogf/gf/v2/frame/g"
)

type (
	GetSystemInitializedReq struct {
		g.Meta `path:"/v1/config/system/initialized" tags:"Config" method:"get" summary:"get system initialized"`
	}
	GetSystemInitializedRes struct {
		Initialized bool `json:"initialized"`
	}
)
