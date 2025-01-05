package account

import (
	"context"

	"gf-user/api/account/user"
	"gf-user/internal/service"
)

func (c *ControllerUser) GetInfo(ctx context.Context, _ *user.GetInfoReq) (res *user.GetInfoRes, err error) {
	ua, err := service.Account().GetUserAccount(ctx)
	if err != nil {
		return
	}

	res = (*user.GetInfoRes)(ua)
	return
}
