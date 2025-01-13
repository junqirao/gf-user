package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerLogin) GetLoginConfig(ctx context.Context, _ *login.GetLoginConfigReq) (res *login.GetLoginConfigRes, err error) {
	cfg := new(model.LoginConfig)
	err = service.Config().Get(ctx, consts.ConfigKeyLogin, cfg)
	res = (*login.GetLoginConfigRes)(cfg)
	return
}
