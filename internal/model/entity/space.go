// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Space is the golang structure for table space.
type Space struct {
	Id          int         `json:"id"          orm:"id"          description:""`                 //
	Name        string      `json:"name"        orm:"name"        description:"space name"`       // space name
	Logo        string      `json:"logo"        orm:"logo"        description:"space logo"`       // space logo
	Owner       string      `json:"owner"       orm:"owner"       description:"owner account id"` // owner account id
	Description string      `json:"description" orm:"description" description:""`                 //
	Profile     string      `json:"profile"     orm:"profile"     description:""`                 //
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"   description:""`                 //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`                 //
}
