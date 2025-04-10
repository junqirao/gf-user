package account

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
)

type (
	GetBindMFAGetQRCodeReq struct {
		g.Meta `path:"/v1/account/mfa/qr" tags:"User" method:"get" summary:"User bind mfa qr-code"`
		model.AuthorizationRequired
	}
	GetBindMFAGetQRCodeRes struct {
		Image string `json:"image"`
	}

	BindMFAReq struct {
		g.Meta `path:"/v1/account/mfa/bind" tags:"User" method:"post" summary:"User bind mfa"`
		model.AuthorizationRequired
		Code string `json:"code"`
	}
	BindMFARes struct{}

	UnBindMFAReq struct {
		g.Meta `path:"/v1/account/mfa/unbind" tags:"User" method:"post" summary:"User unbind mfa"`
		model.AuthorizationRequired
		Password string `json:"password"`
		Nonce    string `json:"nonce"`
		Code     string `json:"code"`
	}
	UnBindMFARes struct{}
)
