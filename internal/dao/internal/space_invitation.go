// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaceInvitationDao is the data access object for the table space_invitation.
type SpaceInvitationDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns SpaceInvitationColumns // columns contains all the column names of Table for convenient usage.
}

// SpaceInvitationColumns defines and stores column names for the table space_invitation.
type SpaceInvitationColumns struct {
	Id        string //
	Space     string //
	From      string //
	Status    string // 0: create, 1: accept, 2: reject, 3: cancel
	Target    string //
	Comment   string //
	CreatedAt string //
}

// spaceInvitationColumns holds the columns for the table space_invitation.
var spaceInvitationColumns = SpaceInvitationColumns{
	Id:        "id",
	Space:     "space",
	From:      "from",
	Status:    "status",
	Target:    "target",
	Comment:   "comment",
	CreatedAt: "created_at",
}

// NewSpaceInvitationDao creates and returns a new DAO object for table data access.
func NewSpaceInvitationDao() *SpaceInvitationDao {
	return &SpaceInvitationDao{
		group:   "default",
		table:   "space_invitation",
		columns: spaceInvitationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaceInvitationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaceInvitationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaceInvitationDao) Columns() SpaceInvitationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaceInvitationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaceInvitationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaceInvitationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
