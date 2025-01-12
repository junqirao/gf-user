package consts

// account
const (
	// status

	AccountStatusNormal = 0
	AccountStatusFrozen = 1

	// type

	AccountTypeNormal = 0
	AccountTypeApp    = 1
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
	// token

	ConfigKeyToken     = "token"
	ConfigStoNameToken = "token_config"

	// mfa

	ConfigKeyMfa     = "mfa"
	ConfigStoNameMfa = "mfa_config"
)
