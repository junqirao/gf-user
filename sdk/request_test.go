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
		Timeout:       5,
	})
)

func TestSDK_ValidateAccessToken(t *testing.T) {
	resp, err := testSDK.ValidateAccessToken(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsInN1YiI6IkVaRHZ3TjBRIiwiYXVkIjpbImNjM2E1MTY3LWJhMTMtNDhkMC1iY2EwLTRjZjU1NTg1MzdhYSJdLCJleHAiOjE3Mzk2MjMyODUsIm5iZiI6MTczOTYxNjA4NSwiaWF0IjoxNzM5NjE2MDg1LCJzaWQiOiIxIiwidWlkIjoiNDMifQ.jLK88tlUwOHTrpQN4a71Z9C3p6mat43LV07mBoIEgJw")
	if err != nil {
		t.Fatal(err)
		return
	}

	s, _ := json.MarshalIndent(resp.Data, "", "  ")
	t.Log(string(s))
}
