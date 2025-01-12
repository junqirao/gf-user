package fn

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gocomponents/kvdb"

	"gf-user/internal/service"
)

func setConfigIfNotExists(ctx context.Context, key string, getFunc func(ctx context.Context) any) (err error) {
	mutex, err := kvdb.NewMutex(ctx, "config_set_if_not_exists")
	if err != nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()

	exist, err := service.Config().Exist(ctx, key)
	if err != nil {
		return
	}
	if exist {
		g.Log().Infof(ctx, "config %s already exists, skip.", key)
		return
	}

	return service.Config().Set(ctx, key, getFunc(ctx))
}
