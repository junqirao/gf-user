package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/internal/service"
)

func (c *ControllerLogin) UserLogout(ctx context.Context, req *login.UserLogoutReq) (res *login.UserLogoutRes, err error) {
	err = service.Account().UserLogout(ctx, req.RefreshToken)
	return
}
