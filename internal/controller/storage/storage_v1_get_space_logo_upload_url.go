package storage

import (
	"context"

	"gf-user/api/storage/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetSpaceLogoUploadUrl(ctx context.Context, _ *v1.GetSpaceLogoUploadUrlReq) (res *v1.GetSpaceLogoUploadUrlRes, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	u, key, err := service.Storage().SignSpaceLogoImageUploadUrl(ctx, token.SpaceId)
	if err != nil {
		return
	}
	res = &v1.GetSpaceLogoUploadUrlRes{
		Url: u,
		Key: key,
	}
	return
}
