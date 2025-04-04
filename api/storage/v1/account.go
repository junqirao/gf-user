package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetAccountAvatarUploadUrlReq struct {
		g.Meta `path:"/v1/storage/account/avatar/upload/url" tags:"Storage" method:"get" summary:"get account avatar upload url"`
		model.AuthorizationRequired
	}
	GetAccountAvatarUploadUrlRes struct {
		Url string `json:"url"`
		Key string `json:"key"`
	}
	GetSpaceLogoUploadUrlReq struct {
		g.Meta `path:"/v1/storage/space/logo/upload/url" tags:"Storage" method:"get" summary:"get space logo upload url"`
		model.AuthorizationRequired
	}
	GetSpaceLogoUploadUrlRes struct {
		Url string `json:"url"`
		Key string `json:"key"`
	}
)
