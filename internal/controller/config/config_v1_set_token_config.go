package config

import (
	"context"

	"gf-user/api/config/v1"
	"gf-user/internal/consts"
	"gf-user/internal/service"
)

func (c *ControllerV1) SetTokenConfig(ctx context.Context, req *v1.SetTokenConfigReq) (res *v1.SetTokenConfigRes, err error) {
	if req.Content != nil && req.Content["token_key"] != nil {
		delete(req.Content, "token_key")
	}
	err = service.Config().Set(ctx, consts.ConfigKeyToken, req.Content)
	return
}
