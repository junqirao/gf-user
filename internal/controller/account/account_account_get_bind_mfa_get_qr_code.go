package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/internal/service"
)

func (c *ControllerAccount) GetBindMFAGetQRCode(ctx context.Context, _ *account.GetBindMFAGetQRCodeReq) (res *account.GetBindMFAGetQRCodeRes, err error) {
	qrCode, err := service.Account().GenerateMFAQRCode(ctx)
	if err != nil {
		return
	}

	res = &account.GetBindMFAGetQRCodeRes{
		Image: qrCode,
	}
	return
}
