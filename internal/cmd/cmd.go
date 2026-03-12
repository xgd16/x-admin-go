package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"x-admin/internal/controller/admin"
	"x-admin/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 全局中间件：JWT 校验，noAuth:"true" 的接口（如登录）跳过
			s.Use(service.Middleware().AuthJwt("admin"))
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().HandlerResponse)
				group.Bind(admin.NewV1())
			})
			s.Run()
			return nil
		},
	}
)
