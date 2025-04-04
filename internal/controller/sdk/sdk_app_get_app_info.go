package sdk

import (
	"context"

	"gf-user/api/sdk/app"
	"gf-user/internal/service"
)

func (c *ControllerApp) GetAppInfo(ctx context.Context, req *app.GetAppInfoReq) (res *app.GetAppInfoRes, err error) {
	info, err := service.App().Info(ctx, req.AppId)
	if err != nil {
		return
	}
	res = (*app.GetAppInfoRes)(info)
	return
}
