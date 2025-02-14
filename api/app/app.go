// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"context"

	"gf-user/api/app/public"
	"gf-user/api/app/v1"
)

type IAppPublic interface {
	GetAppOne(ctx context.Context, req *public.GetAppOneReq) (res *public.GetAppOneRes, err error)
}

type IAppV1 interface {
	CreateApp(ctx context.Context, req *v1.CreateAppReq) (res *v1.CreateAppRes, err error)
	UpdateApp(ctx context.Context, req *v1.UpdateAppReq) (res *v1.UpdateAppRes, err error)
	RemoveApp(ctx context.Context, req *v1.RemoveAppReq) (res *v1.RemoveAppRes, err error)
	ListApp(ctx context.Context, req *v1.ListAppReq) (res *v1.ListAppRes, err error)
	GetAppOne(ctx context.Context, req *v1.GetAppOneReq) (res *v1.GetAppOneRes, err error)
}
