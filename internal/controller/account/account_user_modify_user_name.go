package account

import (
	"context"

	"gf-user/api/account/user"
	"gf-user/internal/service"
)

func (c *ControllerUser) ModifyUserName(ctx context.Context, req *user.ModifyUserNameReq) (res *user.ModifyUserNameRes, err error) {
	err = service.User().ModifyUserName(ctx, req.Name)
	return
}
