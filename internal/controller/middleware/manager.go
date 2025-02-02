package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/service"
)

// MustSuperAdmin space_id = 1 && is manager && code matches
func MustSuperAdmin(r *ghttp.Request) {
	ok, err := service.User().IsSuperAdmin(r.Context())
	if err != nil {
		response.Error(r, err)
		return
	}
	if !ok {
		response.Error(r, response.CodePermissionDeny)
		return
	}

	r.Middleware.Next()
}
