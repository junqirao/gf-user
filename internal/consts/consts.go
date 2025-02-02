package consts

// account
const (
	// status

	AccountStatusNormal = 0
	AccountStatusFrozen = 1

	// type

	AccountTypeNormal = 0
	AccountTypeApp    = 1

	// extra key

	AccountExtraKeyAdminCode = "admin_code"
)

// user
const (
	UserTypeNormal  = 0
	UserTypeManager = 1
)

// token
const (
	DefaultTokenIssuer        = "user-center"
	DefaultAccessTokenExpire  = 60 * 60 * 2
	DefaultRefreshTokenExpire = 60 * 60 * 24 * 30
	DefaultTokenKey           = "dEfaUlTuSerT0k3nK3y"

	CtxKeyTokenInfo = "__token_info"
)

// token from
const (
	TokenFromApp     = "app"
	TokenFromBrowser = "browser"
	TokenFromUnknown = "unknown"
)

// space
const (
	DefaultSpaceId = 1

	SpaceInvitationStatusCreate = 0
	SpaceInvitationStatusAccept = 1
	SpaceInvitationStatusReject = 2
	SpaceInvitationStatusCancel = 3
)

// config
const (
	// system

	ConfigKeySystemInitialized = "system_initialized"

	// token

	ConfigKeyToken     = "token"
	ConfigStoNameToken = "token_config"

	// mfa

	ConfigKeyMfa     = "mfa"
	ConfigStoNameMfa = "mfa_config"
	ConfigKeyLogin   = "login"
)

type LoginMode int8

func (m LoginMode) Name() string {
	switch m {
	case LoginModeMFA:
		return "MFA"
	case LoginModePassword:
		return "Password"
	case LoginModePasswordMFA:
		return "Password+MFA"
	default:
		return "Unknown"
	}
}

const (
	// LoginModePassword only password login
	LoginModePassword LoginMode = 1 << iota
	// LoginModeMFA only mfa login
	LoginModeMFA

	// LoginModePasswordMFA combine LoginModePassword and LoginModeMFA
	LoginModePasswordMFA = LoginModePassword | LoginModeMFA
)
