// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package storage

import (
	"context"

	"gf-user/api/storage/redirect"
	"gf-user/api/storage/v1"
)

type IStorageRedirect interface {
	AccountAvatar(ctx context.Context, req *redirect.AccountAvatarReq) (res *redirect.AccountAvatarRes, err error)
	SpaceLogo(ctx context.Context, req *redirect.SpaceLogoReq) (res *redirect.SpaceLogoRes, err error)
}

type IStorageV1 interface {
	GetAccountAvatarUploadUrl(ctx context.Context, req *v1.GetAccountAvatarUploadUrlReq) (res *v1.GetAccountAvatarUploadUrlRes, err error)
	GetSpaceLogoUploadUrl(ctx context.Context, req *v1.GetSpaceLogoUploadUrlReq) (res *v1.GetSpaceLogoUploadUrlRes, err error)
}
