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
		Get(ctx context.Context, sto string, key string, ptr any) (err error)
		Set(ctx context.Context, sto string, key string, val any) (err error)
		SetIfNotExist(ctx context.Context, name string, key string, val any) (err error)
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
