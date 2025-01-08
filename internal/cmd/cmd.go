package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/junqirao/gocomponents/response"
	"github.com/junqirao/gocomponents/structs"

	"gf-user/internal/controller/account"
	"gf-user/internal/controller/middleware"
	"gf-user/internal/controller/space"
	"gf-user/internal/controller/storage"
)

var (
	tagParser = structs.NewTagParser(structs.WithTagHandlerValueMapping())
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(response.MiddlewareWithDataHandler(func(ctx context.Context, res any) any {
					tagParser.Parse(ctx, res)
					return res
				}))
				// no need authentication
				group.Bind(
					// login & register
					account.NewLogin(),
				)
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.AuthToken)
					group.Bind(
						// storage
						storage.NewV1(),
						// account & user
						account.NewUser(),
						account.NewAccount(),
						// space
						space.NewV1(),
					)
				})
			})
			// redirect
			s.Group("", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.AuthToken)
				group.Bind(
					storage.NewRedirect(),
				)
			})
			s.Run()
			return nil
		},
	}
)
