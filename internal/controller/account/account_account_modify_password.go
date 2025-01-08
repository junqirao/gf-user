package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) ModifyPassword(ctx context.Context, req *account.ModifyPasswordReq) (res *account.ModifyPasswordRes, err error) {
	err = service.Account().ModifyPassword(ctx, req.Old, req.New, req.Nonce)
	return
}
