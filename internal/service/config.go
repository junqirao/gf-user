// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-user/internal/model"
)

type (
	IConfig interface {
		GetTokenConfig(ctx context.Context) (res *model.UserTokenConfig)
		GetMFAConfig(ctx context.Context) (res *model.MFAConfig)
		GetLoginConfig(ctx context.Context) (res *model.LoginConfig)
		Get(ctx context.Context, key string, ptr any) (err error)
		Set(ctx context.Context, key string, val any) (err error)
		Exist(ctx context.Context, key string) (exist bool, err error)
		SystemInitialized(ctx context.Context) (ok bool, err error)
		UpdateSystemInitialized(ctx context.Context) (err error)
	}
)

var (
	localConfig IConfig
)

func Config() IConfig {
	if localConfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return localConfig
}

func RegisterConfig(i IConfig) {
	localConfig = i
}
