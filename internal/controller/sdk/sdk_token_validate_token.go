package sdk

import (
	"context"

	"gf-user/api/sdk/token"
	"gf-user/internal/service"
)

func (c *ControllerToken) ValidateToken(ctx context.Context, req *token.ValidateTokenReq) (res *token.ValidateTokenRes, err error) {
	info, err := service.Token().ValidAccessToken(ctx, req.AccessToken)
	if err != nil {
		return
	}
	res = (*token.ValidateTokenRes)(info)
	return
}
