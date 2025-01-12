// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"gf-user/api/account/account"
	"gf-user/api/account/login"
	"gf-user/api/account/user"
)

type IAccountAccount interface {
	GetBindMFAGetQRCode(ctx context.Context, req *account.GetBindMFAGetQRCodeReq) (res *account.GetBindMFAGetQRCodeRes, err error)
	BindMFA(ctx context.Context, req *account.BindMFAReq) (res *account.BindMFARes, err error)
	ModifyName(ctx context.Context, req *account.ModifyNameReq) (res *account.ModifyNameRes, err error)
	ModifyAvatar(ctx context.Context, req *account.ModifyAvatarReq) (res *account.ModifyAvatarRes, err error)
	ModifyPassword(ctx context.Context, req *account.ModifyPasswordReq) (res *account.ModifyPasswordRes, err error)
}

type IAccountLogin interface {
	UserLogin(ctx context.Context, req *login.UserLoginReq) (res *login.UserLoginRes, err error)
	Register(ctx context.Context, req *login.RegisterReq) (res *login.RegisterRes, err error)
	RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error)
}

type IAccountUser interface {
	GetInfo(ctx context.Context, req *user.GetInfoReq) (res *user.GetInfoRes, err error)
}
