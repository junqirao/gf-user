package config

import (
	"context"

	"gf-user/api/config/public"
	"gf-user/internal/service"
)

func (c *ControllerPublic) GetSystemInitialized(ctx context.Context, req *public.GetSystemInitializedReq) (res *public.GetSystemInitializedRes, err error) {
	initialized, err := service.Config().SystemInitialized(ctx)
	if err != nil {
		return
	}

	res = &public.GetSystemInitializedRes{
		Initialized: initialized,
	}
	return
}
