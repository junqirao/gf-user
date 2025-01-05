package storage

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/junqirao/gocomponents/storage"

	"gf-user/internal/service"
)

func init() {
	service.RegisterStorage(&sStorage{})
}

type sStorage struct {
}

func (s sStorage) SignAvatarImageUploadUrl(ctx context.Context, accountId string) (url, key string, err error) {
	if _, err = service.Account().IsValid(ctx, accountId); err != nil {
		return
	}
	key = grand.S(8)
	url, err = storage.Default().SignPutUrl(ctx, s.getAccountAssetsAvatarKey(accountId, key), 300)
	return
}

func (s sStorage) SignAvatarImageGetUrl(ctx context.Context, accountId, key string) (url string, err error) {
	cacheKey := fmt.Sprintf("account:assets:cache:avatar:%s:%s", accountId, key)
	v, err := g.Redis().Get(ctx, cacheKey)
	if err != nil {
		return
	}
	if u := v.String(); u != "" {
		url = u
		return
	}

	exp := int64(60 * 60 * 24)
	url, err = storage.Default().SignGetUrl(ctx, s.getAccountAssetsAvatarKey(accountId, key), exp, "image/jpeg", "inline;")
	if err != nil {
		return
	}

	if _, err = g.Redis().Set(ctx, cacheKey, url); err != nil {
		return
	}
	_, err = g.Redis().Expire(ctx, cacheKey, exp)
	return
}

func (s sStorage) getAccountAssetsAvatarKey(accountId, key string) string {
	return fmt.Sprintf("account/assets/avatar/%s/%s", accountId, key)
}
