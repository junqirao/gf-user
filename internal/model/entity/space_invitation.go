// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaceInvitation is the golang structure for table space_invitation.
type SpaceInvitation struct {
	Id        int         `json:"id"        orm:"id"         description:""`                                           //
	Space     int         `json:"space"     orm:"space"      description:""`                                           //
	From      string      `json:"from"      orm:"from"       description:""`                                           //
	Status    int         `json:"status"    orm:"status"     description:"0: create, 1: accept, 2: reject, 3: cancel"` // 0: create, 1: accept, 2: reject, 3: cancel
	Target    string      `json:"target"    orm:"target"     description:""`                                           //
	Comment   string      `json:"comment"   orm:"comment"    description:""`                                           //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                                           //
}
