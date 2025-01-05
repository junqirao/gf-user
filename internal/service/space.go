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
	ISpace interface {
		GetSpaceInfo(ctx context.Context, spaceId int64) (space *model.Space, err error)
		CreateSpace(ctx context.Context, in model.CreateSpaceInput) (space *model.Space, err error)
		CreateInvitation(ctx context.Context, in model.CreateSpaceInvitationInput) (err error)
		HandleInvitation(ctx context.Context, id int, op int) (err error)
		MyInvitations(ctx context.Context) (target []*model.SpaceInvitation, source []*model.SpaceInvitation, err error)
	}
)

var (
	localSpace ISpace
)

func Space() ISpace {
	if localSpace == nil {
		panic("implement not found for interface ISpace, forgot register?")
	}
	return localSpace
}

func RegisterSpace(i ISpace) {
	localSpace = i
}
