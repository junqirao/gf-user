package model

type (
	UserTokenConfig struct {
		AccessTokenExpire  int64  `json:"access_token_expire" default:"7200"`
		RefreshTokenExpire int64  `json:"refresh_token_expire" default:"2592000"`
		RefreshTokenLimit  int64  `json:"refresh_token_limit"`
		TokenKey           string `json:"token_key,omitempty" default:"dEfaUlTuSerT0k3nK3y"`
	}
	MFAConfig struct {
		Enable            bool   `json:"enable" default:"false"`
		CodeLength        int    `json:"code_length,omitempty" default:"6"`
		SecretLength      int    `json:"secret_length,omitempty" default:"16"`
		VerifyDiscrepancy int    `json:"verify_discrepancy,omitempty" default:"1"`
		Title             string `json:"title,omitempty" default:"UserCenter"`
		Secret            string `json:"secret,omitempty"`
	}
)
