// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"context"

	"gf-user/api/config/public"
	"gf-user/api/config/v1"
)

type IConfigPublic interface {
	GetSystemInitialized(ctx context.Context, req *public.GetSystemInitializedReq) (res *public.GetSystemInitializedRes, err error)
}

type IConfigV1 interface {
	SetLoginConfig(ctx context.Context, req *v1.SetLoginConfigReq) (res *v1.SetLoginConfigRes, err error)
	GetLoginConfig(ctx context.Context, req *v1.GetLoginConfigReq) (res *v1.GetLoginConfigRes, err error)
	SetMFAConfig(ctx context.Context, req *v1.SetMFAConfigReq) (res *v1.SetMFAConfigRes, err error)
	GetMFAConfig(ctx context.Context, req *v1.GetMFAConfigReq) (res *v1.GetMFAConfigRes, err error)
	SetTokenConfig(ctx context.Context, req *v1.SetTokenConfigReq) (res *v1.SetTokenConfigRes, err error)
	GetTokenConfig(ctx context.Context, req *v1.GetTokenConfigReq) (res *v1.GetTokenConfigRes, err error)
}
