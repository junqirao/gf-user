package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) BindMFA(ctx context.Context, req *account.BindMFAReq) (res *account.BindMFARes, err error) {
	err = service.Account().BindMFA(ctx, req.Code)
	return
}
