package router

import (
	"go-gf-demo/app/api"
	"go-gf-demo/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// s1.SetNameToUriType(ghttp.URI_TYPE_DEFAULT)
	// 	URI_TYPE_DEFAULT  = 0 // （默认）全部转为小写，单词以'-'连接符号连接
	// URI_TYPE_FULLNAME = 1 // 不处理名称，以原有名称构建成URI
	// URI_TYPE_ALLLOWER = 2 // 仅转为小写，单词间不使用连接符号
	// URI_TYPE_CAMEL    = 3 // 采用驼峰命名方式
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
		)
		// 对应的对象，剩下的路由是自动实现的吗？路由声明？
		// 对象注册
		group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(
				service.Middleware.Auth,
			)
			// 对象函数注册
			// 两个group含义是什么
			// 为什么是重复注册了？
			// group.ALL("/user/profile", api.User.Profile)
		})
	})
	// 函数注册
	s.BindHandler("/demo", func(r *ghttp.Request) {
		r.Response.Writeln("go frame demo")
	})
	// namespace??
	s.Domain("local").BindHandler("/demodomain", func(r *ghttp.Request) {
		r.Response.Writeln("localhost go frame demo")
	})
	// 整体是深度优先 /name/list /:name/update /:name/:action  /:name/*any  /:name
	// 最原生
	s.BindHandler("GET:/{class}-{course}/:name/*act", func(r *ghttp.Request) {
		r.Response.Writef(
			"%v %v %v %v", r.Get("class"), r.Get("course"), r.Get("name"), r.Get("act"),
		)
	})
	// :name 命名匹配规则 URI层级必须有值 且到此为止
	// list 精准匹配规则
	// *act 模糊匹配规则 URI层级可以为空
	// {page} 字段匹配规则
	s.BindHandler("/{table}/list/{page}.html", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Router)
	})
	// 绑定路由方法
	// s.BindObjectMethod("/user", api.User, "Show")
	// RESTful对象注册
	// s.BindObjectRest("/object", api.User)
	s.SetPort(8100, 8200)
	// s.Run()
}
