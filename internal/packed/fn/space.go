package fn

import (
	"context"

	"github.com/junqirao/gocomponents/launcher"

	"gf-user/internal/dao"
	"gf-user/internal/model/entity"
	"gf-user/internal/packed"
)

func checkAndInitSpace() *launcher.HookTask {
	return launcher.NewHookTask(
		"check_and_init_space",
		func(ctx context.Context) (err error) {
			count, err := dao.Space.Ctx(ctx).Where(dao.Space.Columns().Id, packed.DefaultSpaceId).Count()
			if err != nil {
				return
			}
			if count != 0 {
				return
			}

			_, err = dao.Space.Ctx(ctx).Insert(entity.Space{
				Id:          packed.DefaultSpaceId,
				Name:        "Default",
				Owner:       "system",
				Description: "Default",
				Profile:     "{}",
			})
			return
		},
	)
}
