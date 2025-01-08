package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) ModifyName(ctx context.Context, req *account.ModifyNameReq) (res *account.ModifyNameRes, err error) {
	err = service.Account().ModifyName(ctx, req.Name)
	return
}
