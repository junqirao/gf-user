package model

import (
	"encoding/json"

	"github.com/gogf/gf/v2/os/gtime"

	"gf-user/internal/model/entity"
)

type AppInfo struct {
	Id           string         `json:"id"` // app uuid (app_id)
	Name         any            `json:"name"`
	Space        int            `json:"space"`
	Descriptions string         `json:"descriptions"`
	Profile      map[string]any `json:"profile"`
	ExpiredAt    *gtime.Time    `json:"expiredAt"`
	CreatedAt    *gtime.Time    `json:"createdAt"`
}

func NewAppInfo(a *entity.App) *AppInfo {
	profile := make(map[string]any)
	if a.Profile != "" {
		_ = json.Unmarshal([]byte(a.Profile), &profile)
	}
	return &AppInfo{
		Id:           a.Id,
		Name:         a.Name,
		Space:        a.Space,
		Descriptions: a.Descriptions,
		Profile:      profile,
		ExpiredAt:    a.ExpiredAt,
		CreatedAt:    a.CreatedAt,
	}
}

type CreateAppInput struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Profile     map[string]interface{} `json:"profile"`
	ExpiredAt   *gtime.Time            `json:"expired_at"`
}

type CreateAppOutput struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type ValidateAppInput struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Nonce     string `json:"nonce"`
}

type UpdateAppInput struct {
	AppId       string                 `json:"app_id"`
	Name        *string                `json:"name"`
	Description *string                `json:"description"`
	Profile     map[string]interface{} `json:"profile"`
	ExpiredAt   *gtime.Time            `json:"expired_at"`
}
