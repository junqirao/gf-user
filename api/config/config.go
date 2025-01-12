// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"context"

	"gf-user/api/config/v1"
)

type IConfigV1 interface {
	SetTokenConfig(ctx context.Context, req *v1.SetTokenConfigReq) (res *v1.SetTokenConfigRes, err error)
	GetTokenConfig(ctx context.Context, req *v1.GetTokenConfigReq) (res *v1.GetTokenConfigRes, err error)
	SetMFAConfig(ctx context.Context, req *v1.SetMFAConfigReq) (res *v1.SetMFAConfigRes, err error)
	GetMFAConfig(ctx context.Context, req *v1.GetMFAConfigReq) (res *v1.GetMFAConfigRes, err error)
}
