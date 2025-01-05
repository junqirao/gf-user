// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaceDao is the data access object for the table space.
type SpaceDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns SpaceColumns // columns contains all the column names of Table for convenient usage.
}

// SpaceColumns defines and stores column names for the table space.
type SpaceColumns struct {
	Id          string //
	Name        string // space name
	Logo        string // space logo
	Owner       string // owner account id
	Description string //
	Profile     string //
	UpdateAt    string //
	CreatedAt   string //
}

// spaceColumns holds the columns for the table space.
var spaceColumns = SpaceColumns{
	Id:          "id",
	Name:        "name",
	Logo:        "logo",
	Owner:       "owner",
	Description: "description",
	Profile:     "profile",
	UpdateAt:    "update_at",
	CreatedAt:   "created_at",
}

// NewSpaceDao creates and returns a new DAO object for table data access.
func NewSpaceDao() *SpaceDao {
	return &SpaceDao{
		group:   "default",
		table:   "space",
		columns: spaceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaceDao) Columns() SpaceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
