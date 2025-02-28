package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/service"
)

func (c *ControllerLogin) GenerateAppToken(ctx context.Context, req *login.GenerateAppTokenReq) (res *login.GenerateAppTokenRes, err error) {
	loginInfo, err := service.Account().GenerateAppToken(ctx, req.AppId, req.Space, req.RefreshToken)
	if err != nil {
		return
	}

	res = (*login.GenerateAppTokenRes)(loginInfo)
	return
}
