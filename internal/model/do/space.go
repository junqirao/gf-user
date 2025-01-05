// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Space is the golang structure of table space for DAO operations like Where/Data.
type Space struct {
	g.Meta      `orm:"table:space, do:true"`
	Id          interface{} //
	Name        interface{} // space name
	Logo        interface{} // space logo
	Owner       interface{} // owner account id
	Description interface{} //
	Profile     interface{} //
	UpdateAt    *gtime.Time //
	CreatedAt   *gtime.Time //
}
