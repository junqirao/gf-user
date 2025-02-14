package app

import (
	"context"

	"gf-user/api/app/v1"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (c *ControllerV1) ListApp(ctx context.Context, _ *v1.ListAppReq) (res *v1.ListAppRes, err error) {
	isMgr, err := service.User().IsSpaceManager(ctx)
	if err != nil {
		return
	}
	if !isMgr {
		err = code.ErrNotSpaceManager
		return
	}

	out, err := service.App().List(ctx)
	if err != nil {
		return
	}

	res = &v1.ListAppRes{
		List: out,
	}
	return
}
