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
}

type IStorageV1 interface {
	GetAccountAvatarUploadUrl(ctx context.Context, req *v1.GetAccountAvatarUploadUrlReq) (res *v1.GetAccountAvatarUploadUrlRes, err error)
}
