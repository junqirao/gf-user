package config

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gocomponents/kvdb"
	"github.com/junqirao/gocomponents/structs"

	"gf-user/internal/model"
	"gf-user/internal/packed"
	"gf-user/internal/service"
)

var (
	once = sync.Once{}
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
	err := s.getOne(ctx, s.getName(packed.ConfigStoNameToken), packed.ConfigKeyToken, res)
	if err != nil {
		g.Log().Warningf(ctx, "failed to get token config: %s", err.Error())
		return
	}
	return
}

func (s sConfig) Get(ctx context.Context, sto, key string, ptr any) (err error) {
	return s.getOne(ctx, s.getName(sto), key, ptr)
}

func (s sConfig) Set(ctx context.Context, sto, key string, val any) (err error) {
	return s.set(ctx, s.getName(sto), key, val)
}

func (s sConfig) SetIfNotExist(ctx context.Context, name, key string, val any) (err error) {
	var res []*kvdb.KV
	res, err = kvdb.Storages.GetStorage(name).Get(ctx, key)
	switch {
	case err == nil:
		if len(res) > 0 {
			return
		}
	case errors.Is(err, kvdb.ErrStorageNotFound):
		err = nil
	default:
		return
	}

	return kvdb.Storages.GetStorage(name).Set(ctx, key, val)
}

func (s sConfig) getName(topic string) string {
	return fmt.Sprintf("user_servic_config_%s", topic)
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
	return kvdb.Storages.GetStorage(name).Set(ctx, key, val)
}
