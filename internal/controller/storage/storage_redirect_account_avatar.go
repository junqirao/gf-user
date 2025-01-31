package storage

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	"gf-user/api/storage/redirect"
	"gf-user/internal/service"
)

func (c *ControllerRedirect) AccountAvatar(ctx context.Context, req *redirect.AccountAvatarReq) (res *redirect.AccountAvatarRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	token := service.Token().GetTokenInfoFromCtx(ctx)
	url, err := service.Storage().SignAvatarImageGetUrl(ctx, token.AccountId, req.Key)
	if err != nil {
		return
	}
	r.Response.RedirectTo(url)
	r.Response.Flush()
	return
}
