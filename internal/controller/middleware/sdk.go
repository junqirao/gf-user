package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
	"github.com/junqirao/gf-user/sdk"
)

func AuthSdk(r *ghttp.Request) {
	appId, appSecret, nonce, err := sdk.DecodeAuthenticationStr(r.Header.Get("Authorization"))
	if err != nil {
		response.Error(r, code.ErrInvalidAppToken)
		return
	}
	err = service.App().Auth(r.Context(), &model.ValidateAppInput{
		AppId:     appId,
		AppSecret: appSecret,
		Nonce:     nonce,
	})
	if err != nil {
		response.Error(r, err)
		return
	}
	r.Middleware.Next()
}
