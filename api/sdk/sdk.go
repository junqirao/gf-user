// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sdk

import (
	"context"

	"gf-user/api/sdk/token"
)

type ISdkToken interface {
	ValidateToken(ctx context.Context, req *token.ValidateTokenReq) (res *token.ValidateTokenRes, err error)
}
