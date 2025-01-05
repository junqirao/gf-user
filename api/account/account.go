// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"gf-user/api/account/login"
	"gf-user/api/account/user"
)

type IAccountLogin interface {
	UserLogin(ctx context.Context, req *login.UserLoginReq) (res *login.UserLoginRes, err error)
	Register(ctx context.Context, req *login.RegisterReq) (res *login.RegisterRes, err error)
	RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error)
}

type IAccountUser interface {
	GetInfo(ctx context.Context, req *user.GetInfoReq) (res *user.GetInfoRes, err error)
}
