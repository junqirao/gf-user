package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) UnBindMFA(ctx context.Context, req *account.UnBindMFAReq) (res *account.UnBindMFARes, err error) {
	err = service.Account().UnbindMFA(ctx, req.Password, req.Nonce, req.Code)
	return
}
