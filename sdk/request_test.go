package sdk

import (
	"context"
	"encoding/json"
	"testing"
)

var (
	testSDK = NewSDK(&Config{
		Address:       "http://127.0.0.1:8000",
		AppId:         "7f1e5b4a-3696-442b-8f4d-d434ab941a59",
		AppSecret:     "oh8fZ3XeXZs0DcgCCUsoT26O7vWAlxldXfZhF0M0qwiqZi71",
		SkipTLSVerify: true,
		Timeout:       5000,
	})
)

func TestSDK_ValidateAccessToken(t *testing.T) {
	resp, err := testSDK.ValidateAccessToken(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsInN1YiI6Ilk0VXFrRk55IiwiYXVkIjpbImNjM2E1MTY3LWJhMTMtNDhkMC1iY2EwLTRjZjU1NTg1MzdhYSJdLCJleHAiOjE3NDA4NTI3MzUsIm5iZiI6MTc0MDg0NTUzNSwiaWF0IjoxNzQwODQ1NTM1LCJzaWQiOiIxIiwidWlkIjoiNDMiLCJhcHAiOiI3ZjFlNWI0YS0zNjk2LTQ0MmItOGY0ZC1kNDM0YWI5NDFhNTkifQ.cKmUBjKESQCQelo6UC8lgqsNh5DqWlvIUh_hCgBdpZY")
	if err != nil {
		t.Fatal(err)
		return
	}

	s, _ := json.MarshalIndent(resp.Data, "", "  ")
	t.Log(string(s))
}
