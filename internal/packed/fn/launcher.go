package fn

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/junqirao/gocomponents/launcher"
	"github.com/junqirao/gocomponents/storage"
	"github.com/junqirao/gocomponents/structs"

	"gf-user/internal/consts"
	"gf-user/internal/packed"
	"gf-user/internal/packed/ipgeo"
)

func BeforeTasks() []*launcher.HookTask {
	return []*launcher.HookTask{
		launcher.NewHookTask("setup_go_frame_module", func(ctx context.Context) error {
			setupPProf(ctx)
			setupDebugMode(ctx)
			return nil
		}),
		checkAndInitSpace(),
		launcher.NewHookTask("load_struct_mapping", func(ctx context.Context) error {
			return structs.LoadMappingFromEmbed(ctx, packed.Embed)
		}),
		launcher.NewHookTask("init_storage_module", func(ctx context.Context) error {
			storage.MustInit(ctx)
			return nil
		}),
		launcher.NewHookTask("init_user_token_config", func(ctx context.Context) error {
			return setConfigIfNotExists(ctx, consts.ConfigKeyToken, func(ctx context.Context) any {
				cfg := map[string]any{
					"token_key": grand.S(16),
				}
				g.Log().Info(ctx, "token config generate token key.")
				return cfg
			})
		}),
		launcher.NewHookTask("init_mfa_config", func(ctx context.Context) error {
			return setConfigIfNotExists(ctx, consts.ConfigKeyMfa, func(ctx context.Context) any {
				cfg := map[string]any{
					"secret": grand.S(16),
				}
				g.Log().Info(ctx, "mfa config generate secret.")
				return cfg
			})
		}),
		launcher.NewHookTask("init_ipgeo_module", func(ctx context.Context) error {
			ipgeo.Init(ctx)
			return nil
		}),
	}
}
