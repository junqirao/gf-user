// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaceInvitation is the golang structure for table space_invitation.
type SpaceInvitation struct {
	Id        int         `json:"id"        orm:"id"         description:""` //
	Space     int         `json:"space"     orm:"space"      description:""` //
	From      string      `json:"from"      orm:"from"       description:""` //
	Target    string      `json:"target"    orm:"target"     description:""` //
	Comment   string      `json:"comment"   orm:"comment"    description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
}
