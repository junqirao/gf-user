package redirect

import (
	"github.com/gogf/gf/v2/frame/g"
)

type (
	AccountAvatarReq struct {
		g.Meta `path:"/v1/storage/account/avatar" tags:"Storage" method:"get" summary:"redirect get account avatar"`
		// model.AuthenticationRequired
		Key string `json:"key"`
	}
	AccountAvatarRes struct {
	}
)
