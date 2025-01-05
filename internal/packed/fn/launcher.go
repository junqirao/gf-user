package fn

import (
	"context"

	"github.com/junqirao/gocomponents/launcher"
	"github.com/junqirao/gocomponents/storage"
	"github.com/junqirao/gocomponents/structs"

	"gf-user/internal/packed"
)

func BeforeTasks() []*launcher.HookTask {
	return []*launcher.HookTask{
		checkAndInitSpace(),
		launcher.NewHookTask("load_struct_mapping", func(ctx context.Context) error {
			return structs.LoadMappingFromEmbed(ctx, packed.Embed)
		}),
		launcher.NewHookTask("init_storage_module", func(ctx context.Context) error {
			storage.MustInit(ctx)
			return nil
		}),
	}
}
