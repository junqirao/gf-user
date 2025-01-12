package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerAccount) ModifyPassword(ctx context.Context, req *account.ModifyPasswordReq) (res *account.ModifyPasswordRes, err error) {
	err = service.Account().ModifyPassword(ctx, &model.AccountModifyPasswordInput{
		MFACodeRequired: req.MFACodeRequired,
		Old:             req.Old,
		New:             req.New,
		Nonce:           req.Nonce,
	})
	return
}
