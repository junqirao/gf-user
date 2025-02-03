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
	if err = service.Config().Get(ctx, consts.ConfigKeyLogin, cfg); err != nil {
		return
	}
	mfa := new(model.MFAConfig)
	err = service.Config().Get(ctx, consts.ConfigKeyMfa, mfa)
	res = &login.GetLoginConfigRes{
		LoginConfig:   *cfg,
		MFAEnabled:    mfa.Enable,
		MFACodeLength: mfa.CodeLength,
	}
	return
}
