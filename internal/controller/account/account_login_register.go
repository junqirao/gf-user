package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerLogin) Register(ctx context.Context, req *login.RegisterReq) (res *login.RegisterRes, err error) {
	out, err := service.Account().Register(ctx, &model.AccountRegisterInput{
		Account:  req.Account,
		Password: req.Password,
		Type:     consts.AccountTypeNormal,
		Status:   consts.AccountStatusNormal,
		Name:     req.Name,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Extra:    make(map[string]any),
	})
	if err != nil {
		return
	}

	res = (*login.RegisterRes)(out)
	return
}
