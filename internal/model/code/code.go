package code

import (
	"net/http"

	"github.com/junqirao/gocomponents/response"
)

var (
	ErrAccountNotExist = response.NewCode(1000, "account not exist", http.StatusNotFound)
	ErrAccountLocked   = response.NewCode(1001, "account frozen", http.StatusLocked)
	ErrAccountPassword = response.NewCode(1002, "account password error", http.StatusUnauthorized)
	ErrUserNotExist    = response.NewCode(1100, "user not exist", http.StatusNotFound)
)

// token
var (
	ErrInvalidToken = response.NewCode(2000, "invalid token", http.StatusUnauthorized)
)
