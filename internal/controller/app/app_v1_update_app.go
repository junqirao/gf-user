package app

import (
	"context"

	"gf-user/api/app/v1"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (c *ControllerV1) UpdateApp(ctx context.Context, req *v1.UpdateAppReq) (res *v1.UpdateAppRes, err error) {
	isMgr, err := service.User().IsSpaceManager(ctx)
	if err != nil {
		return
	}
	if !isMgr {
		err = code.ErrNotSpaceManager
		return
	}
	err = service.App().Update(ctx, &req.UpdateAppInput)
	return
}
