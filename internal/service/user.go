// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-user/internal/model/do"
)

type (
	IUser interface {
		GetUserFromToken(ctx context.Context, spaceId int64) (usr *do.User, err error)
		GetUserByAccountId(ctx context.Context, accountId string, spaceId int64) (usr *do.User, err error)
		Exist(ctx context.Context, accountId string, spaceId int64) (exist bool, err error)
		CreateSpaceUser(ctx context.Context, account *do.Account, spaceId int64) (usr *do.User, err error)
		IsSpaceManager(ctx context.Context) (ok bool, err error)
		IsSuperAdmin(ctx context.Context) (ok bool, err error)
		ModifyUserName(ctx context.Context, name string) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
