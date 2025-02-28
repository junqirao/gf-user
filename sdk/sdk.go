package sdk

import (
	"crypto/tls"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (r *Response[T]) parse(resp *gclient.Response) (err error) {
	if err = gjson.Unmarshal(resp.ReadAll(), &r); err != nil {
		return
	}
	if resp.StatusCode > 399 {
		if r.Message == "" {
			err = errors.New(http.StatusText(resp.StatusCode))
		} else {
			err = errors.New(r.Message)
		}
	}
	return
}

type SDK struct {
	cfg *Config
	cli *gclient.Client
}

func NewSDK(cfg *Config) *SDK {
	cli := gclient.New()
	transport := &http.Transport{
		MaxIdleConns:    10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: time.Second * 90,
	}
	if cfg.SkipTLSVerify {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	cli.Transport = transport
	cli.SetTimeout(time.Millisecond * time.Duration(cfg.Timeout))
	return &SDK{
		cfg: cfg,
		cli: cli,
	}
}

func (s SDK) client(noAuth ...bool) *gclient.Client {
	cli := s.cli.Clone()
	if !(len(noAuth) > 0 && noAuth[0]) {
		cli.SetHeader("Authorization", GenerateAuthenticationStr(s.cfg.AppId, s.cfg.AppSecret))
		cli.SetHeader(HeaderKeyAppid, s.cfg.AppId)
	}
	return cli
}

func (s SDK) url(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return s.cfg.Address + path
}
