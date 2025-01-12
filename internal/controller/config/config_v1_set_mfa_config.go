package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/consts"
	"gf-user/internal/service"
)

func (c *ControllerV1) SetMFAConfig(ctx context.Context, req *v1.SetMFAConfigReq) (res *v1.SetMFAConfigRes, err error) {
	if req.Content != nil && req.Content["secret"] != nil {
		delete(req.Content, "secret")
	}
	err = service.Config().Set(ctx, consts.ConfigKeyMfa, req.Content)
	return
}
