// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-user/internal/model"
	"gf-user/internal/model/do"
)

type (
	IAccount interface {
		Register(ctx context.Context, in *model.AccountRegisterInput) (out *model.UserAccount, err error)
		UserLogin(ctx context.Context, in *model.AccountLoginInput) (out *model.UserAccountLoginInfo, err error)
		UserLogout(ctx context.Context, refreshToken string) (err error)
		IsValid(ctx context.Context, accountId string) (acc *do.Account, err error)
		RefreshToken(ctx context.Context, spaceId int64, refreshToken string) (res *model.UserAccountLoginInfo, err error)
		GenerateAppToken(ctx context.Context, appId string, spaceId int64, refreshToken string) (res *model.UserAccountLoginInfo, err error)
		GetUserAccount(ctx context.Context, spaceId ...int64) (ua *model.UserAccount, err error)
		GetAccount(ctx context.Context, account string) (acc *do.Account, err error)
		GetAccountById(ctx context.Context, id string) (acc *model.Account, err error)
		GetAccountByIds(ctx context.Context, id []string) (acs []*model.Account, err error)
		Exists(ctx context.Context, account string) (exists bool, err error)
		RegisterAdministrator(ctx context.Context, in *model.AccountRegisterInput) (out *model.UserAccount, err error)
		GenerateMFAQRCode(ctx context.Context) (qrCode string, err error)
		BindMFA(ctx context.Context, mfaCode string) (err error)
		VerifyMFACode(ctx context.Context, account *do.Account, mfaCode string) (err error)
		UnbindMFA(ctx context.Context, pwd string, nonce string, cod string) (err error)
		ModifyName(ctx context.Context, name string) (err error)
		ModifyAvatar(ctx context.Context, avatar string) (err error)
		ModifyPassword(ctx context.Context, in *model.AccountModifyPasswordInput) (err error)
	}
)

var (
	localAccount IAccount
)

func Account() IAccount {
	if localAccount == nil {
		panic("implement not found for interface IAccount, forgot register?")
	}
	return localAccount
}

func RegisterAccount(i IAccount) {
	localAccount = i
}
