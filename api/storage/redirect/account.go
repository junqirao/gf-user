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
	SpaceLogoReq struct {
		g.Meta `path:"/v1/storage/space/logo" tags:"Storage" method:"get" summary:"redirect get space logo"`
		// model.AuthenticationRequired
		Key     string `json:"key"`
		SpaceId int64  `json:"space_id"`
	}
	SpaceLogoRes struct {
	}
)
