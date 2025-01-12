package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) SetTokenConfig(ctx context.Context, req *v1.SetTokenConfigReq) (res *v1.SetTokenConfigRes, err error) {
	err = service.Config().SetTokenConfig(ctx, req.Content)
	return
}
