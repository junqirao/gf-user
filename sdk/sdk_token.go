package sdk

import (
	"context"
	"net/http"
)

func (s SDK) ValidateAccessToken(ctx context.Context, accessToken string) (info *Response[*TokenInfo], err error) {
	resp, err := s.client().SetHeader(HeaderKeyAppToken, accessToken).DoRequest(ctx, http.MethodPost, s.url("/sdk/token/validate"), http.NoBody)
	if err != nil {
		return
	}
	info = new(Response[*TokenInfo])
	err = info.parse(resp)
	return
}
