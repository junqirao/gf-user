package model

type (
	AppAuthorizationRequired struct {
		AuthorizationRequired
		AppId    string `p:"X-App-Id" in:"header" v:"required"`
		AppToken string `p:"X-App-Token" in:"header" v:"required"`
	}
)
