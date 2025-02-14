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
	IApp interface {
		Create(ctx context.Context, in *model.CreateAppInput) (out *model.CreateAppOutput, err error)
		List(ctx context.Context) (infos []*model.AppInfo, err error)
		Info(ctx context.Context, id string) (info *model.AppInfo, err error)
		Update(ctx context.Context, in *model.UpdateAppInput) (err error)
		Auth(ctx context.Context, in *model.ValidateAppInput) (err error)
		Remove(ctx context.Context, id string) (err error)
	}
)

var (
	localApp IApp
)

func App() IApp {
	if localApp == nil {
		panic("implement not found for interface IApp, forgot register?")
	}
	return localApp
}

func RegisterApp(i IApp) {
	localApp = i
}
