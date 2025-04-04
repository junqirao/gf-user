package sdk

import (
	"github.com/junqirao/gf-user/define"
)

const (
	HeaderKeyAppid    = "X-App-Id"
	HeaderKeyAppToken = "X-App-Token"
)

type (
	TokenInfo            = define.TokenInfo
	Space                = define.Space
	AccountBrief         = define.AccountBrief
	Account              = define.Account
	UserAccount          = define.UserAccount
	UserAccountLoginInfo = define.UserAccountLoginInfo
	UserInfo             = define.UserInfo
	AppInfo              = define.AppInfo
)
