package space

import (
	"context"

	"gf-user/api/space/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) MyInvitations(ctx context.Context, _ *v1.MyInvitationsReq) (res *v1.MyInvitationsRes, err error) {
	target, source, err := service.Space().MyInvitations(ctx)
	if err != nil {
		return
	}
	res = &v1.MyInvitationsRes{
		Target: target,
		Source: source,
	}
	return
}
