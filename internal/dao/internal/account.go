// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountDao is the data access object for the table account.
type AccountDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns AccountColumns // columns contains all the column names of Table for convenient usage.
}

// AccountColumns defines and stores column names for the table account.
type AccountColumns struct {
	Id        string // account uuid
	Account   string // unique account
	Password  string // password hash
	Type      string // 0: normal, 1: app
	Status    string // 0: normal, 1: frozen
	Name      string //
	Email     string //
	Avatar    string //
	CreatedAt string //
	UpdateAt  string //
	Extra     string // extra
}

// accountColumns holds the columns for the table account.
var accountColumns = AccountColumns{
	Id:        "id",
	Account:   "account",
	Password:  "password",
	Type:      "type",
	Status:    "status",
	Name:      "name",
	Email:     "email",
	Avatar:    "avatar",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
	Extra:     "extra",
}

// NewAccountDao creates and returns a new DAO object for table data access.
func NewAccountDao() *AccountDao {
	return &AccountDao{
		group:   "default",
		table:   "account",
		columns: accountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AccountDao) Columns() AccountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
