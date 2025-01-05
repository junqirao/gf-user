// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaceInvitation is the golang structure of table space_invitation for DAO operations like Where/Data.
type SpaceInvitation struct {
	g.Meta    `orm:"table:space_invitation, do:true"`
	Id        interface{} //
	Space     interface{} //
	From      interface{} //
	Status    interface{} // 0: create, 1: accept, 2: reject, 3: cancel
	Target    interface{} //
	Comment   interface{} //
	CreatedAt *gtime.Time //
}
