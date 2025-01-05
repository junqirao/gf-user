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
		IsValid(ctx context.Context, accountId string) (acc *do.Account, err error)
		RefreshToken(ctx context.Context, spaceId int64, refreshToken string) (res *model.UserAccountLoginInfo, err error)
		GetUserAccount(ctx context.Context, spaceId ...int64) (ua *model.UserAccount, err error)
		GetAccount(ctx context.Context, account string) (acc *do.Account, err error)
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
