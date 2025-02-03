// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package space

import (
	"context"

	"gf-user/api/space/v1"
)

type ISpaceV1 interface {
	CreateSpace(ctx context.Context, req *v1.CreateSpaceReq) (res *v1.CreateSpaceRes, err error)
	CreateSpaceInvitation(ctx context.Context, req *v1.CreateSpaceInvitationReq) (res *v1.CreateSpaceInvitationRes, err error)
	HandleSpaceInvitation(ctx context.Context, req *v1.HandleSpaceInvitationReq) (res *v1.HandleSpaceInvitationRes, err error)
	MyInvitations(ctx context.Context, req *v1.MyInvitationsReq) (res *v1.MyInvitationsRes, err error)
	GetAccountSpaces(ctx context.Context, req *v1.GetAccountSpacesReq) (res *v1.GetAccountSpacesRes, err error)
}
