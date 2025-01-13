package fn

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	onceLoadPProf = sync.Once{}
)

func enableFunctionWithConfig(ctx context.Context, pattern string, def bool, do func(), elseDo ...func()) {
	enable := def
	v, err := g.Cfg().Get(ctx, pattern)
	if err == nil {
		enable = v.Bool()
	}

	if !enable {
		if len(elseDo) > 0 && elseDo[0] != nil {
			elseDo[0]()
		}
		return
	}
	do()
}

func setupPProf(ctx context.Context, serverName ...string) {
	onceLoadPProf.Do(func() {
		name := ghttp.DefaultServerName
		if len(serverName) > 0 {
			name = serverName[0]
		}
		enableFunctionWithConfig(ctx, "server.pprof", true, func() {
			prefix := ""
			v, err := g.Cfg().Get(ctx, "server.pprof_prefix")
			if err == nil {
				prefix = v.String()
			}
			g.Server(name).EnablePProf(prefix)
		})
	})
}

func setupDebugMode(ctx context.Context) {
	enableFunctionWithConfig(ctx, "server.debug", true,
		// if enable
		func() {
			g.Log().SetDebug(true)
			g.Server().SetDumpRouterMap(true)
			g.Log().Info(ctx, "debug mode enabled.")
		},
		// else
		func() {
			g.Log().SetDebug(false)
			g.Server().SetDumpRouterMap(false)
		})
}
