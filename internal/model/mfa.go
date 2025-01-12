package model

type MFACodeRequired struct {
	MFACode string `json:"X-MFA-Code" in:"header"`
}
