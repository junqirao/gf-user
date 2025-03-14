package sdk

import (
	"context"
	"net/http"
)

func (s SDK) GetUserInfo(ctx context.Context, accessToken string) (info *Response[*UserAccount], err error) {
	resp, err := s.client().SetHeader(HeaderKeyAppToken, accessToken).DoRequest(ctx, http.MethodGet, s.url("/sdk/user/userinfo"), http.NoBody)
	if err != nil {
		return
	}
	info = new(Response[*UserAccount])
	err = info.parse(resp)
	return
}
