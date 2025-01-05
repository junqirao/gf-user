package middleware

import (
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/consts"
	"gf-user/internal/service"
)

func AuthToken(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		token = r.Get("access_token").String()
	}
	if token == "" {
		response.Error(r, response.CodeInvalidParameter.WithDetail("Authorization is required"))
		return
	}
	info, err := service.Token().ValidAccessToken(r.Context(), strings.TrimPrefix(token, "Bearer "))
	if err != nil {
		response.Error(r, err)
		return
	}
	r.SetCtxVar(consts.CtxKeyTokenInfo, info)
	r.Middleware.Next()
}
