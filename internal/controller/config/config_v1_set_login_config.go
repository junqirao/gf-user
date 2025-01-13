package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/consts"
	"gf-user/internal/service"
)

func (c *ControllerV1) SetLoginConfig(ctx context.Context, req *v1.SetLoginConfigReq) (res *v1.SetLoginConfigRes, err error) {
	err = service.Config().Set(ctx, consts.ConfigKeyLogin, req.Content)
	return
}
