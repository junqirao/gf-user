package model

import (
	"gf-user/internal/consts"
)

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
	LoginConfig struct {
		LoginMode       consts.LoginMode   `json:"login_mode" default:"1"`
		EnableRegister  bool               `json:"enable_register" default:"false"`
		RememberAccount bool               `json:"remember_account" default:"false"`
		MaximumFail     int                `json:"maximum_fail" default:"5"`
		Display         LoginDisplayConfig `json:"display"`
	}
	LoginDisplayConfig struct {
		Notice string `json:"notice"`
	}
	SystemInitializeFlag struct {
		Initialized bool `json:"initialized"`
	}
)
