package app

import (
	"context"

	"gf-user/api/app/v1"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (c *ControllerV1) CreateApp(ctx context.Context, req *v1.CreateAppReq) (res *v1.CreateAppRes, err error) {
	isMgr, err := service.User().IsSpaceManager(ctx)
	if err != nil {
		return
	}
	if !isMgr {
		err = code.ErrNotSpaceManager
		return
	}

	out, err := service.App().Create(ctx, &req.CreateAppInput)
	if err != nil {
		return
	}

	res = (*v1.CreateAppRes)(out)
	return
}
