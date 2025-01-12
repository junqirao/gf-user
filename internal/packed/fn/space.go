package fn

import (
	"context"

	"github.com/junqirao/gocomponents/kvdb"
	"github.com/junqirao/gocomponents/launcher"

	"gf-user/internal/consts"
	"gf-user/internal/dao"
	"gf-user/internal/model/entity"
)

func checkAndInitSpace() *launcher.HookTask {
	return launcher.NewHookTask(
		"check_and_init_space",
		func(ctx context.Context) (err error) {
			mutex, err := kvdb.NewMutex(ctx, "user_service_init_space")
			if err != nil {
				return err
			}
			mutex.Lock()
			defer mutex.Unlock()

			count, err := dao.Space.Ctx(ctx).Where(dao.Space.Columns().Id, consts.DefaultSpaceId).Count()
			if err != nil {
				return
			}
			if count != 0 {
				return
			}

			_, err = dao.Space.Ctx(ctx).Insert(entity.Space{
				Id:          consts.DefaultSpaceId,
				Name:        "Default",
				Owner:       "system",
				Description: "Default",
				Profile:     "{}",
			})
			return
		},
	)
}
