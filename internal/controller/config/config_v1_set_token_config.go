package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/packed"
	"gf-user/internal/service"
)

func (c *ControllerV1) SetTokenConfig(ctx context.Context, req *v1.SetTokenConfigReq) (res *v1.SetTokenConfigRes, err error) {
	err = service.Config().Set(ctx, packed.ConfigStoNameToken, packed.ConfigKeyToken, req.Content)
	return
}
