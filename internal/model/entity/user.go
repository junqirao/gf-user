// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"        orm:"id"         description:"user id"`               // user id
	Account   string      `json:"account"   orm:"account"    description:"account id"`            // account id
	Space     int         `json:"space"     orm:"space"      description:"space id"`              // space id
	Type      int         `json:"type"      orm:"type"       description:"0: normal, 1: manager"` // 0: normal, 1: manager
	Name      string      `json:"name"      orm:"name"       description:""`                      //
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:""`                      //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                      //
}
