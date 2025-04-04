package sdk

import (
	"context"
	"net/http"
)

func (s SDK) GetAppInfo(ctx context.Context, accessToken string) (info *Response[*AppInfo], err error) {
	resp, err := s.client().SetHeader(HeaderKeyAppToken, accessToken).DoRequest(ctx, http.MethodGet, s.url("/sdk/app/info"), http.NoBody)
	if err != nil {
		return
	}
	info = new(Response[*AppInfo])
	err = info.parse(resp)
	return
}

func (s SDK) GetAppInfoPublic(ctx context.Context) (info *Response[*AppInfo], err error) {
	resp, err := s.client().DoRequest(ctx, http.MethodGet, s.url("/public/app/"+s.cfg.AppId), http.NoBody)
	if err != nil {
		return
	}
	info = new(Response[*AppInfo])
	err = info.parse(resp)
	return
}
