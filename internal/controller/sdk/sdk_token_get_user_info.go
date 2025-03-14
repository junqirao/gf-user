package sdk

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"

	"gf-user/api/sdk/token"
	"gf-user/internal/consts"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
	"github.com/junqirao/gf-user/sdk"
)

func (c *ControllerToken) GetUserInfo(ctx context.Context, _ *token.GetUserInfoReq) (res *token.GetUserInfoRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		tokenStr = r.Header.Get(sdk.HeaderKeyAppToken)
	)

	if tokenStr == "" {
		err = code.ErrTokenRequired
		return
	}
	info, err := service.Token().ValidAccessToken(ctx, strings.TrimPrefix(tokenStr, "Bearer "))
	if err != nil {
		return
	}

	ua, err := service.Account().GetUserAccount(context.WithValue(ctx, consts.CtxKeyTokenInfo, info))
	if err != nil {
		return
	}

	res = (*token.GetUserInfoRes)(ua)
	return
}
