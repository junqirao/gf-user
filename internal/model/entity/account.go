// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure for table account.
type Account struct {
	Id        string      `json:"id"        orm:"id"         description:"account uuid"`         // account uuid
	Account   string      `json:"account"   orm:"account"    description:"unique account"`       // unique account
	Password  string      `json:"password"  orm:"password"   description:"password hash"`        // password hash
	Type      int         `json:"type"      orm:"type"       description:"0: normal, 1: app"`    // 0: normal, 1: app
	Status    int         `json:"status"    orm:"status"     description:"0: normal, 1: frozen"` // 0: normal, 1: frozen
	Name      string      `json:"name"      orm:"name"       description:""`                     //
	Email     string      `json:"email"     orm:"email"      description:""`                     //
	Avatar    string      `json:"avatar"    orm:"avatar"     description:""`                     //
	Mfa       []byte      `json:"mfa"       orm:"mfa"        description:""`                     //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                     //
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:""`                     //
	Extra     string      `json:"extra"     orm:"extra"      description:"extra"`                // extra
}
