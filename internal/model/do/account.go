// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure of table account for DAO operations like Where/Data.
type Account struct {
	g.Meta    `orm:"table:account, do:true"`
	Id        interface{} // account uuid
	Account   interface{} // unique account
	Password  interface{} // password hash
	Type      interface{} // 0: normal, 1: app
	Status    interface{} // 0: normal, 1: frozen
	Name      interface{} //
	Email     interface{} //
	Avatar    interface{} //
	CreatedAt *gtime.Time //
	UpdateAt  *gtime.Time //
	Extra     interface{} // extra
}
