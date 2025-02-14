// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// App is the golang structure for table app.
type App struct {
	Id           string      `json:"id"           orm:"id"           description:"app uuid (app_id)"` // app uuid (app_id)
	Secret       string      `json:"secret"       orm:"secret"       description:"secret md5"`        // secret md5
	Name         string      `json:"name"         orm:"name"         description:""`                  //
	Space        int         `json:"space"        orm:"space"        description:""`                  //
	Descriptions string      `json:"descriptions" orm:"descriptions" description:""`                  //
	Profile      string      `json:"profile"      orm:"profile"      description:""`                  //
	ExpiredAt    *gtime.Time `json:"expiredAt"    orm:"expired_at"   description:""`                  //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"   description:""`                  //
}
