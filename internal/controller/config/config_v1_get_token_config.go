package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetTokenConfig(ctx context.Context, _ *v1.GetTokenConfigReq) (res *v1.GetTokenConfigRes, err error) {
	cfg := service.Config().GetTokenConfig(ctx)
	// desensitization
	cfg.TokenKey = ""
	res = (*v1.GetTokenConfigRes)(cfg)
	return
}
