package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetLoginConfig(ctx context.Context, _ *v1.GetLoginConfigReq) (res *v1.GetLoginConfigRes, err error) {
	cfg := new(model.LoginConfig)
	err = service.Config().Get(ctx, consts.ConfigKeyLogin, cfg)
	res = (*v1.GetLoginConfigRes)(cfg)
	return
}
