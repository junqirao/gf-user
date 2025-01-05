package fn

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gocomponents/updater"

	"gf-user/internal/packed"
)

func UpdaterFuncInfos(ctx context.Context) []*updater.FuncInfo {
	return append(updater.SQLFuncFromEmbedFS(ctx, g.DB(), packed.Embed)) // fill your function
}
