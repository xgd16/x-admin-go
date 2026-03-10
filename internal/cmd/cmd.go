package cmd

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"x-admin/internal/controller/auth"
	"x-admin/internal/controller/user"
	"x-admin/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 全局中间件：对非登录接口进行 JWT 校验
			s.Use(func(r *ghttp.Request) {
				if strings.HasSuffix(r.URL.Path, "/auth/login") {
					r.Middleware.Next()
					return
				}
				middleware.AuthJwt(r)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					auth.NewV1(),
					user.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
