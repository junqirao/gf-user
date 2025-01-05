package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	CreateSpaceReq struct {
		g.Meta `path:"/v1/space" tags:"Space" method:"put" summary:"Create space"`
		model.AuthorizationRequired
		Name        string `json:"name"`
		Logo        string `json:"logo"`
		Description string `json:"description"`
	}
	CreateSpaceRes           model.Space
	CreateSpaceInvitationReq struct {
		g.Meta `path:"/v1/space/invitation" tags:"Space" method:"put" summary:"Create space invitation"`
		model.AuthorizationRequired
		Account string `json:"account"`
		Comment string `json:"comment"`
	}
	CreateSpaceInvitationRes struct{}
	HandleSpaceInvitationReq struct {
		g.Meta `path:"/v1/space/invitation/handle" tags:"Space" method:"post" summary:"Accept space invitation"`
		model.AuthorizationRequired
		Id        int `json:"id"`
		Operation int `json:"operation"`
	}
	HandleSpaceInvitationRes struct{}
	MyInvitationsReq         struct {
		g.Meta `path:"/v1/space/invitation/my" tags:"Space" method:"get" summary:"My space invitations"`
		model.AuthorizationRequired
	}
	MyInvitationsRes struct {
		Source []*model.SpaceInvitation `json:"source"`
		Target []*model.SpaceInvitation `json:"target"`
	}
)
