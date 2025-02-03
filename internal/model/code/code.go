package code

import (
	"net/http"

	"github.com/junqirao/gocomponents/response"
)

// user & account
var (
	ErrAccountNotExist       = response.NewCode(1000, "account not exist", http.StatusNotFound)
	ErrAccountLocked         = response.NewCode(1001, "account frozen", http.StatusLocked)
	ErrAccountPassword       = response.NewCode(1002, "account password error", http.StatusUnauthorized)
	ErrAccountMfaAlreadyBind = response.NewCode(1003, "account mfa already bind", http.StatusBadRequest)
	ErrAccountMfaCode        = response.NewCode(1004, "account mfa code error", http.StatusBadRequest)
	ErrAccountMfaNotBind     = response.NewCode(1005, "account mfa not bind", http.StatusBadRequest)
	ErrAccountAlreadyExist   = response.NewCode(1006, "account already exist", http.StatusConflict)
	ErrAdminAlreadyExist     = response.NewCode(1007, "administrator already exist", http.StatusConflict)
	ErrAdminCodeMismatch     = response.NewCode(1008, "admin code mismatch", http.StatusForbidden)
	ErrUserNotExist          = response.NewCode(1100, "user not exist", http.StatusNotFound)
)

// token
var (
	ErrInvalidToken         = response.NewCode(1200, "invalid token", http.StatusUnauthorized)
	ErrTokenRequired        = response.NewCode(1201, "token required", http.StatusBadRequest)
	ErrRefreshTokenNotFound = response.NewCode(1202, "invalid token", http.StatusBadRequest)
)

// space
var (
	ErrNotSpaceManager = response.NewCode(1300, "not space manager", http.StatusForbidden)
)
