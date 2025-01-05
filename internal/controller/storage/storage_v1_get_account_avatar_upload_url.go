package storage

import (
	"context"

	"gf-user/api/storage/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) GetAccountAvatarUploadUrl(ctx context.Context, req *v1.GetAccountAvatarUploadUrlReq) (res *v1.GetAccountAvatarUploadUrlRes, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	u, key, err := service.Storage().SignAvatarImageUploadUrl(ctx, token.AccountId)
	if err != nil {
		return
	}
	res = &v1.GetAccountAvatarUploadUrlRes{
		Url: u,
		Key: key,
	}
	return
}
