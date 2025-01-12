package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/packed"
	"gf-user/internal/service"
)

// MustSuperAdmin space_id = 1 && is manager
func MustSuperAdmin(r *ghttp.Request) {
	ok, err := service.User().IsSpaceManager(r.Context())
	if err != nil {
		response.Error(r, err)
		return
	}
	if !ok {
		response.Error(r, response.CodePermissionDeny)
		return
	}

	info := service.Token().GetTokenInfoFromCtx(r.Context())
	if info.SpaceId != packed.DefaultSpaceId {
		response.Error(r, response.CodePermissionDeny.WithDetail("invalid space"))
		return
	}
	r.Middleware.Next()
}
