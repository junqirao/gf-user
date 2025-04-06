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
	"gf-user/internal/dao"
	"gf-user/internal/model"
	"gf-user/internal/model/do"
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

func (s sConfig) SystemInitialized(ctx context.Context) (ok bool, err error) {
	val, err := kvdb.Storages.GetStorage(stoName).Get(ctx, consts.ConfigKeySystemInitialized)
	switch {
	case errors.Is(err, kvdb.ErrStorageNotFound):
		err = s.updateSystemInitialized(ctx)
	case len(val) == 0:
		err = s.updateSystemInitialized(ctx)
	case err == nil:
		get := func() (ok bool, err error) {
			ptr := new(model.SystemInitializeFlag)
			if err = val[0].Value.Struct(&ptr); err != nil {
				return
			}
			ok = ptr.Initialized
			return
		}
		ok, err = get()
		if !ok {
			if err = s.updateSystemInitialized(ctx); err != nil {
				return
			}
			ok, err = get()
		}
		return
	default:
		return
	}

	ok = err == nil
	return
}

func (s sConfig) UpdateSystemInitialized(ctx context.Context) (err error) {
	return s.updateSystemInitialized(ctx)
}

func (s sConfig) updateSystemInitialized(ctx context.Context) (err error) {
	mutex, err := kvdb.NewMutex(ctx, "get_system_status")
	if err != nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()

	var users []*do.User
	err = dao.User.Ctx(ctx).
		Where(g.Map{
			dao.User.Columns().Space: consts.DefaultSpaceId,
			dao.User.Columns().Type:  consts.UserTypeManager,
		}).
		Scan(&users)
	if err != nil {
		return
	}
	if len(users) == 0 {
		err = s.set(ctx, stoName, consts.ConfigKeySystemInitialized, &model.SystemInitializeFlag{Initialized: false})
		return
	}
	var acc []string
	for _, user := range users {
		acc = append(acc, gconv.String(user.Account))
	}
	var accounts []*do.Account
	err = dao.Account.Ctx(ctx).WhereIn(dao.Account.Columns().Id, acc).Scan(&accounts)
	if err != nil {
		return
	}
	cod := g.Cfg().MustGet(ctx, "admin.code").String()
	for _, account := range accounts {
		if cod == gconv.MapStrStr(account.Extra)[consts.AccountExtraKeyAdminCode] {
			err = s.set(ctx, stoName, consts.ConfigKeySystemInitialized, &model.SystemInitializeFlag{Initialized: true})
			return
		}
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
