// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} // user id
	Account   interface{} // account id
	Space     interface{} // space id
	Type      interface{} // 0: normal, 1: manager
	Name      interface{} //
	UpdateAt  *gtime.Time //
	CreatedAt *gtime.Time //
}
