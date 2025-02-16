package sdk

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/util/grand"
)

// GenerateAuthenticationStr
// |md5(app_id)|md5(md5(app_secret)+nonce)|nonce|signature
func GenerateAuthenticationStr(appId, appSecret string) (code string) {
	nonce := grand.S(8)
	var parts []string
	parts = append(parts, appId)
	parts = append(parts, gmd5.MustEncrypt(gmd5.MustEncrypt(appSecret)+nonce))
	parts = append(parts, nonce)
	code = gbase64.EncodeString(strings.Join(parts, "."))
	return
}

func DecodeAuthenticationStr(code string) (appId, appSecret, nonce string, err error) {
	decoded, err := gbase64.DecodeToString(code)
	if err != nil {
		return
	}
	parts := strings.Split(decoded, ".")
	if len(parts) != 3 {
		err = errors.New("invalid code")
		return
	}

	appId = parts[0]
	appSecret = parts[1]
	nonce = parts[2]
	return
}
