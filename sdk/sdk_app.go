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
