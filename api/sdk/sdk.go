// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sdk

import (
	"context"

	"gf-user/api/sdk/app"
	"gf-user/api/sdk/token"
)

type ISdkApp interface {
	GetAppInfo(ctx context.Context, req *app.GetAppInfoReq) (res *app.GetAppInfoRes, err error)
}

type ISdkToken interface {
	ValidateAppToken(ctx context.Context, req *token.ValidateAppTokenReq) (res *token.ValidateAppTokenRes, err error)
	GetUserInfo(ctx context.Context, req *token.GetUserInfoReq) (res *token.GetUserInfoRes, err error)
}
