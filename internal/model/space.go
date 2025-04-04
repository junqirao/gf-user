package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/junqirao/gf-user/define"
)

type Space = define.Space

type SpaceInvitation struct {
	Id        any           `json:"id"`
	Space     any           `json:"space"`
	From      *AccountBrief `json:"from"`
	To        *AccountBrief `json:"to"`
	Comment   any           `json:"comment"`
	CreatedAt *gtime.Time   `json:"created_at"`
}

type CreateSpaceInput struct {
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}

type CreateSpaceInvitationInput struct {
	TargetAccount string `json:"target_account"`
	Comment       string `json:"comment"`
}
