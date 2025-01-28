package login

import (
	"github.com/gogf/gf/v2/frame/g"
)

type (
	CheckExistsReq struct {
		g.Meta  `path:"/v1/account/exists" tags:"User" method:"get" summary:"Account exists"`
		Account string `json:"account"`
	}
	CheckExistsRes struct {
		Exists bool `json:"exists"`
	}
)
