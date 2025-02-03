package space

import (
	"context"

	"gf-user/api/space/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetAccountSpaces(ctx context.Context, _ *v1.GetAccountSpacesReq) (res *v1.GetAccountSpacesRes, err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	res = &v1.GetAccountSpacesRes{}
	res.List, err = service.Space().GetSpaceList(ctx, tokenInfo.AccountId)
	return
}
