package space

import (
	"context"

	"gf-user/api/space/v1"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

func (c *ControllerV1) CreateSpace(ctx context.Context, req *v1.CreateSpaceReq) (res *v1.CreateSpaceRes, err error) {
	space, err := service.Space().CreateSpace(ctx, model.CreateSpaceInput{
		Name:        req.Name,
		Logo:        req.Logo,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	res = (*v1.CreateSpaceRes)(space)
	return
}
