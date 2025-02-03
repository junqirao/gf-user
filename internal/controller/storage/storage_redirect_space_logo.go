package storage

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	"gf-user/api/storage/redirect"
	"gf-user/internal/service"
)

func (c *ControllerRedirect) SpaceLogo(ctx context.Context, req *redirect.SpaceLogoReq) (res *redirect.SpaceLogoRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	sid := req.SpaceId
	if sid <= 0 {
		sid = service.Token().GetTokenInfoFromCtx(ctx).SpaceId
	}
	url, err := service.Storage().SignSpaceLogoImageGetUrl(ctx, sid, req.Key)
	if err != nil {
		return
	}
	r.Response.RedirectTo(url)
	r.Response.Flush()
	return
}
