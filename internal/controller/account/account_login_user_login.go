package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerLogin) UserLogin(ctx context.Context, req *login.UserLoginReq) (res *login.UserLoginRes, err error) {
	loginInfo, err := service.Account().UserLogin(ctx, &model.AccountLoginInput{
		Account:  req.Account,
		Password: req.Password,
		Nonce:    req.Nonce,
		From:     req.From,
	})
	if err != nil {
		return
	}

	res = (*login.UserLoginRes)(loginInfo)
	return
}
