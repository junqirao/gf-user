// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// App is the golang structure of table app for DAO operations like Where/Data.
type App struct {
	g.Meta       `orm:"table:app, do:true"`
	Id           interface{} // app uuid (app_id)
	Secret       interface{} // secret md5
	Name         interface{} //
	Space        interface{} //
	Descriptions interface{} //
	Profile      interface{} //
	ExpiredAt    *gtime.Time //
	CreatedAt    *gtime.Time //
}
