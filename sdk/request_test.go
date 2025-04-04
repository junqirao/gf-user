package sdk

import (
	"context"
	"encoding/json"
	"testing"
)

var (
	testSDK = NewSDK(&Config{
		Address:       "http://127.0.0.1:8001",
		AppId:         "7f1e5b4a-3696-442b-8f4d-d434ab941a59",
		AppSecret:     "oh8fZ3XeXZs0DcgCCUsoT26O7vWAlxldXfZhF0M0qwiqZi71",
		SkipTLSVerify: true,
		Timeout:       5000,
	})
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsInN1YiI6Ik52MFlaQVJWIiwiYXVkIjpbImNjM2E1MTY3LWJhMTMtNDhkMC1iY2EwLTRjZjU1NTg1MzdhYSJdLCJleHAiOjE3NDM3NjE0ODksIm5iZiI6MTc0Mzc1NDI4OSwiaWF0IjoxNzQzNzU0Mjg5LCJzaWQiOiI3IiwidWlkIjoiNDQiLCJhcHAiOiI3ZjFlNWI0YS0zNjk2LTQ0MmItOGY0ZC1kNDM0YWI5NDFhNTkifQ.S9dLKs4_IpEgUI8AkbnIfvfu-JduY4DO_rtZDP19vTE"
)

func TestSDK_ValidateAccessToken(t *testing.T) {
	resp, err := testSDK.ValidateAccessToken(context.Background(), token)
	if err != nil {
		t.Fatal(err)
		return
	}

	s, _ := json.MarshalIndent(resp.Data, "", "  ")
	t.Log(string(s))
}

func TestSDK_GetUserInfo(t *testing.T) {
	resp, err := testSDK.GetUserInfo(context.Background(), token)
	if err != nil {
		t.Fatal(err)
		return
	}

	s, _ := json.MarshalIndent(resp, "", "  ")
	t.Log(string(s))
}

func TestSDK_GetAppInfo(t *testing.T) {
	resp, err := testSDK.GetAppInfo(context.Background(), token)
	if err != nil {
		t.Fatal(err)
		return
	}

	s, _ := json.MarshalIndent(resp, "", "  ")
	t.Log(string(s))
}

func TestSDK_GetAppInfoPublic(t *testing.T) {
	resp, err := testSDK.GetAppInfoPublic(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}

	s, _ := json.MarshalIndent(resp, "", "  ")
	t.Log(string(s))
}
