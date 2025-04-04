// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"gf-user/internal/dao/internal"
)

// internalSpaceInvitationDao is an internal type for wrapping the internal DAO implementation.
type internalSpaceInvitationDao = *internal.SpaceInvitationDao

// spaceInvitationDao is the data access object for the table space_invitation.
// You can define custom methods on it to extend its functionality as needed.
type spaceInvitationDao struct {
	internalSpaceInvitationDao
}

var (
	// SpaceInvitation is a globally accessible object for table space_invitation operations.
	SpaceInvitation = spaceInvitationDao{
		internal.NewSpaceInvitationDao(),
	}
)

// Add your custom methods and functionality below.
