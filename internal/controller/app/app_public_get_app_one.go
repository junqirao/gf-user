package app

import (
	"context"

	"gf-user/api/app/public"
	"gf-user/internal/service"
)

func (c *ControllerPublic) GetAppOne(ctx context.Context, req *public.GetAppOneReq) (res *public.GetAppOneRes, err error) {
	info, err := service.App().Info(ctx, req.AppId)
	if err != nil {
		return
	}

	res = (*public.GetAppOneRes)(info)
	return
}
