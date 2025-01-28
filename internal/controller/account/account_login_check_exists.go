package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/service"
)

func (c *ControllerLogin) CheckExists(ctx context.Context, req *login.CheckExistsReq) (res *login.CheckExistsRes, err error) {
	exists, err := service.Account().Exists(ctx, req.Account)
	if err != nil {
		return
	}
	res = &login.CheckExistsRes{
		Exists: exists,
	}
	return
}
