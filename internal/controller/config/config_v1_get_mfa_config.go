package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetMFAConfig(ctx context.Context, _ *v1.GetMFAConfigReq) (res *v1.GetMFAConfigRes, err error) {
	cfg := new(model.MFAConfig)
	err = service.Config().Get(ctx, consts.ConfigKeyMfa, cfg)
	res = (*v1.GetMFAConfigRes)(cfg)
	return
}
