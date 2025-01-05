package space

import (
	"context"

	"gf-user/api/space/v1"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerV1) CreateSpaceInvitation(ctx context.Context, req *v1.CreateSpaceInvitationReq) (_ *v1.CreateSpaceInvitationRes, err error) {
	err = service.Space().CreateInvitation(ctx, model.CreateSpaceInvitationInput{
		TargetAccount: req.Account,
		Comment:       req.Comment,
	})
	return
}
