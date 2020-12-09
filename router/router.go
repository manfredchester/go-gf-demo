package router

import (
	"go-gf-demo/app/api"
	"go-gf-demo/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
		)
		group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(
				service.Middleware.Auth,
			)
			group.ALL("/user/profile", api.User.Profile)
		})
	})
	s.Run()
}
