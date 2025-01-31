package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetTokenConfig(ctx context.Context, _ *v1.GetTokenConfigReq) (res *v1.GetTokenConfigRes, err error) {
	cfg := new(model.UserTokenConfig)
	err = service.Config().Get(ctx, consts.ConfigKeyToken, cfg)
	// desensitization
	cfg.TokenKey = ""
	res = (*v1.GetTokenConfigRes)(cfg)
	return
}
