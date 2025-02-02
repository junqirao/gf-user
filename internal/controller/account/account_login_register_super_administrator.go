package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerLogin) RegisterSuperAdministrator(ctx context.Context, req *login.RegisterSuperAdministratorReq) (res *login.RegisterSuperAdministratorRes, err error) {
	out, err := service.Account().RegisterAdministrator(ctx, &model.AccountRegisterInput{
		Account:  req.Account,
		Password: req.Password,
		Name:     "Administrator",
		Type:     consts.AccountTypeNormal,
		Status:   consts.AccountStatusNormal,
		Extra: map[string]any{
			consts.AccountExtraKeyAdminCode: req.Code,
			"nonce":                         req.Nonce,
		},
		Administrator: true,
	})
	if err != nil {
		return
	}

	res = (*login.RegisterSuperAdministratorRes)(out)
	return
}
