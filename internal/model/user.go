package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	UserInfo struct {
		Id     interface{} `json:"id"`      // user id
		Space  interface{} `json:"space"`   // space id
		Type   interface{} `json:"type"`    // 0: normal, 1: manager
		Name   interface{} `json:"name"`    //
		JoinAt *gtime.Time `json:"join_at"` // join time
	}
)
