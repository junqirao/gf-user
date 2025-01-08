package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) ModifyAvatar(ctx context.Context, req *account.ModifyAvatarReq) (res *account.ModifyAvatarRes, err error) {
	err = service.Account().ModifyAvatar(ctx, req.Avatar)
	return
}
