package main

import (
	_ "gf-user/internal/logic"

	"context"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/junqirao/gocomponents/kvdb"
	"github.com/junqirao/gocomponents/launcher"
	"github.com/junqirao/gocomponents/updater"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"gf-user/internal/cmd"
	"gf-user/internal/packed/fn"
)

func main() {
	launcher.Launch(
		func(ctx context.Context) {
			cmd.Main.Run(ctx)
		},
		launcher.WithBeforeTasks(fn.BeforeTasks()...),
		launcher.WithConcurrencyUpdater(
			updater.NewKVDatabaseAdaptor(kvdb.Raw),
			fn.UpdaterFuncInfos(gctx.GetInitCtx())...,
		),
	)
}
