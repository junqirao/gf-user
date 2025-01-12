package v1

import (
	"gf-user/internal/model"
)

type SetConfig struct {
	model.AuthorizationRequired
	Content map[string]any `json:"content"`
}
