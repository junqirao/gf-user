package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) GetTokenDetailList(ctx context.Context, req *account.GetTokenDetailListReq) (res *account.GetTokenDetailListRes, err error) {
	details, err := service.Token().ListUserRefreshTokenDetails(ctx, req.Locale)
	if err != nil {
		return
	}
	res = &account.GetTokenDetailListRes{
		List: details,
	}
	return
}
