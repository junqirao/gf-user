package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Space struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Logo        string         `json:"logo"`
	LogoKey     string         `json:"logo_key"`
	IsOwner     bool           `json:"is_owner"`
	Description string         `json:"description"`
	Profile     map[string]any `json:"profile"`
	UpdateAt    *gtime.Time    `json:"update_at"`
	CreatedAt   *gtime.Time    `json:"created_at"`
}

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
