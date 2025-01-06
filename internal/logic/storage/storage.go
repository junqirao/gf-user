package storage

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/junqirao/gocomponents/storage"

	"gf-user/internal/model/code"
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

func (s sStorage) SignSpaceLogoImageUploadUrl(ctx context.Context, spaceId int64) (url, key string, err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	if _, err = service.Account().IsValid(ctx, tokenInfo.AccountId); err != nil {
		return
	}
	ok, err := service.User().IsSpaceManager(ctx)
	if err != nil {
		return
	}
	if !ok {
		err = code.ErrNotSpaceManager
		return
	}

	key = grand.S(8)
	url, err = storage.Default().SignPutUrl(ctx, s.getSpaceAssetsLogoKey(spaceId, key), 300)
	return
}

func (s sStorage) SignAvatarImageGetUrl(ctx context.Context, accountId, key string) (url string, err error) {
	return s.signImageGetUrl(ctx,
		// cache key
		fmt.Sprintf("account:assets:cache:avatar:%s:%s", accountId, key),
		// storage key
		s.getAccountAssetsAvatarKey(accountId, key),
		60*60*24,
	)
}

func (s sStorage) SignSpaceLogoImageGetUrl(ctx context.Context, spaceId int64, key string) (url string, err error) {
	return s.signImageGetUrl(ctx,
		// cache key
		fmt.Sprintf("space:assets:cache:logo:%v:%s", spaceId, key),
		// storage key
		s.getSpaceAssetsLogoKey(spaceId, key),
		60*60*24*7,
	)
}

func (s sStorage) signImageGetUrl(ctx context.Context, cacheKey, storageKey string, expire int64) (url string, err error) {
	v, err := g.Redis().Get(ctx, cacheKey)
	if err != nil {
		return
	}
	if u := v.String(); u != "" {
		url = u
		return
	}

	url, err = storage.Default().SignGetUrl(ctx, storageKey, expire, "image/jpeg", "inline;")
	if err != nil {
		return
	}

	if _, err = g.Redis().Set(ctx, cacheKey, url); err != nil {
		return
	}
	_, err = g.Redis().Expire(ctx, cacheKey, expire)
	return
}

func (s sStorage) getAccountAssetsAvatarKey(accountId, key string) string {
	return fmt.Sprintf("account/assets/avatar/%s/%s", accountId, key)
}

func (s sStorage) getSpaceAssetsLogoKey(spaceId int64, key string) string {
	return fmt.Sprintf("space/assets/logo/%v/%s", spaceId, key)
}
