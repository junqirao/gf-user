package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/service"
)

func (c *ControllerLogin) RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error) {
	loginInfo, err := service.Account().RefreshToken(ctx, req.Space, req.RefreshToken)
	if err != nil {
		return
	}

	res = (*login.RefreshTokenRes)(loginInfo)
	return
}
