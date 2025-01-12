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
}
