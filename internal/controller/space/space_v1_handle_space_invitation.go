package space

import (
	"context"

	"gf-user/api/space/v1"
	"gf-user/internal/service"
)

func (c *ControllerV1) HandleSpaceInvitation(ctx context.Context, req *v1.HandleSpaceInvitationReq) (_ *v1.HandleSpaceInvitationRes, err error) {
	err = service.Space().HandleInvitation(ctx, req.Id, req.Operation)
	return
}
