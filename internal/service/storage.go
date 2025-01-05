// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IStorage interface {
		SignAvatarImageUploadUrl(ctx context.Context, accountId string) (url string, key string, err error)
		SignAvatarImageGetUrl(ctx context.Context, accountId string, key string) (url string, err error)
	}
)

var (
	localStorage IStorage
)

func Storage() IStorage {
	if localStorage == nil {
		panic("implement not found for interface IStorage, forgot register?")
	}
	return localStorage
}

func RegisterStorage(i IStorage) {
	localStorage = i
}
