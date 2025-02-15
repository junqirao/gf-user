package sdk

type Config struct {
	Address       string `json:"address"`
	AppId         string `json:"app_id"`
	AppSecret     string `json:"app_secret"`
	SkipTLSVerify bool   `json:"skip_tls_verify"`
	Timeout       int    `json:"timeout"` // ms
}
