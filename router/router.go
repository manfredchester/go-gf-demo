package router

import (
	"go-gf-demo/app/api"
	"go-gf-demo/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
		)
		// 对应的对象，剩下的路由是自动实现的吗？路由声明？
		group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(
				service.Middleware.Auth,
			)
			// 两个group含义是什么
			group.ALL("/user/profile", api.User.Profile)
		})
	})
	s.Run()
}
