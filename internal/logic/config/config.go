package config

import (
	"context"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gocomponents/kvdb"
	"github.com/junqirao/gocomponents/structs"

	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/service"
)

var (
	once = sync.Once{}
)

const (
	stoName = "user_service_config"
)

// MustInit explicitly, must init before any other service
func MustInit(_ context.Context) {
	once.Do(func() {
		service.RegisterConfig(&sConfig{
			parser: structs.NewTagParser(structs.WithTagHandlerDefaultVal()),
		})
	})
}

type sConfig struct {
	parser *structs.TagParser
}

func (s sConfig) GetTokenConfig(ctx context.Context) (res *model.UserTokenConfig) {
	res = new(model.UserTokenConfig)
	err := s.getOne(ctx, stoName, consts.ConfigKeyToken, res)
	if err != nil {
		g.Log().Warningf(ctx, "failed to get token config: %s", err.Error())
		return
	}
	return
}

func (s sConfig) GetMFAConfig(ctx context.Context) (res *model.MFAConfig) {
	res = new(model.MFAConfig)
	err := s.getOne(ctx, stoName, consts.ConfigKeyMfa, res)
	if err != nil {
		g.Log().Warningf(ctx, "failed to get mfa config: %s", err.Error())
		return
	}
	return
}

func (s sConfig) GetLoginConfig(ctx context.Context) (res *model.LoginConfig) {
	res = new(model.LoginConfig)
	err := s.getOne(ctx, stoName, consts.ConfigKeyLogin, res)
	if err != nil {
		g.Log().Warningf(ctx, "failed to get login config: %s", err.Error())
		return
	}
	return
}

func (s sConfig) Get(ctx context.Context, key string, ptr any) (err error) {
	return s.getOne(ctx, stoName, key, ptr)
}

func (s sConfig) Set(ctx context.Context, key string, val any) (err error) {
	return s.set(ctx, stoName, key, val)
}

func (s sConfig) Exist(ctx context.Context, key string) (exist bool, err error) {
	var res []*kvdb.KV
	res, err = kvdb.Storages.GetStorage(stoName).Get(ctx, key)
	switch {
	case err == nil:
		exist = len(res) > 0
	case errors.Is(err, kvdb.ErrStorageNotFound):
		err = nil
	default:
	}
	return
}

func (s sConfig) getOne(ctx context.Context, name, key string, ptr any, def ...any) (err error) {
	defer func() {
		if ptr == nil {
			if len(def) > 0 && def[0] != nil {
				ptr = def[0]
			}
		}
		s.parser.Parse(ctx, ptr)
	}()
	val, err := kvdb.Storages.GetStorage(name).Get(ctx, key)
	if err != nil || len(val) == 0 {
		if errors.Is(err, kvdb.ErrStorageNotFound) {
			err = nil
		}
		return
	}
	err = val[0].Value.Struct(&ptr)
	return
}

func (s sConfig) set(ctx context.Context, name, key string, val any) (err error) {
	sto := kvdb.Storages.GetStorage(name)
	kvs, err := sto.Get(ctx, key)
	switch {
	case err == nil:
	case errors.Is(err, kvdb.ErrStorageNotFound):
		err = nil
	default:
		return
	}

	if len(kvs) > 0 {
		curr := kvs[0].Value.Map()
		m := gconv.Map(val)
		for k, v := range m {
			curr[k] = v
		}
		val = curr
	}
	return sto.Set(ctx, key, val)
}
