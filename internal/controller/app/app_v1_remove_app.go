package app

import (
	"context"

	"gf-user/api/app/v1"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (c *ControllerV1) RemoveApp(ctx context.Context, req *v1.RemoveAppReq) (res *v1.RemoveAppRes, err error) {
	isMgr, err := service.User().IsSpaceManager(ctx)
	if err != nil {
		return
	}
	if !isMgr {
		err = code.ErrNotSpaceManager
		return
	}
	err = service.App().Remove(ctx, req.AppId)
	return
}
