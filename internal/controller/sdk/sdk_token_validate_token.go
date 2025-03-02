package sdk

import (
	"context"

	"gf-user/api/sdk/token"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (c *ControllerToken) ValidateAppToken(ctx context.Context, req *token.ValidateAppTokenReq) (res *token.ValidateAppTokenRes, err error) {
	info, err := service.Token().ValidAccessToken(ctx, req.AppToken)
	if err != nil {
		return
	}
	if info.AppId == "" || info.AppId != req.AppId {
		err = code.ErrInvalidAppId.WithDetail(info.AppId)
		return
	}
	res = (*token.ValidateAppTokenRes)(info)
	return
}
