package app

import (
	"context"

	"gf-user/api/app/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetAppOne(ctx context.Context, req *v1.GetAppOneReq) (res *v1.GetAppOneRes, err error) {
	info, err := service.App().Info(ctx, req.AppId)
	if err != nil {
		return
	}

	res = (*v1.GetAppOneRes)(info)
	return
}
