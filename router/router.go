package router

import (
	"go-gf-demo/app/api"
	"go-gf-demo/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

func init() {
	s1 := g.Server("s1")
	// URI_TYPE_DEFAULT  = 0 // （默认）全部转为小写，单词以'-'连接符号连接
	// URI_TYPE_FULLNAME = 1 // 不处理名称，以原有名称构建成URI
	// URI_TYPE_ALLLOWER = 2 // 仅转为小写，单词间不使用连接符号
	// URI_TYPE_CAMEL    = 3 // 采用驼峰命名方式
	s1.SetNameToUriType(ghttp.URI_TYPE_DEFAULT)
	s1.SetPort(8100)
	// 分组路由注册方式
	s1.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
		)
		// 对应的对象，剩下的路由是自动实现的吗？自动路由声明？
		group.ALL("/user", api.User)                      // 对象注册
		group.Group("/", func(group *ghttp.RouterGroup) { // 两个group含义是层级注册
			group.Middleware(
				service.Middleware.Auth,
			)
			// 为什么是重复注册了？
			// group.ALL("/user/profile", api.User.Profile) // 对象函数注册
		})
	})

	s2 := g.Server("s2")
	s2.SetNameToUriType(ghttp.URI_TYPE_DEFAULT)
	s2.SetPort(8200)

	s2.BindHandler("/demo1", func(r *ghttp.Request) {
		r.Response.Writeln("go frame demo") // 函数注册
	})
	// namespace??
	s2.Domain("local").BindHandler("/demo1domain", func(r *ghttp.Request) {
		r.Response.Writeln("localhost go frame demo")
	})
	// 整体是深度优先 /name/list /:name/update /:name/:action  /:name/*any  /:name
	// :name 命名匹配规则 URI层级必须有值 且到此为止
	// list 精准匹配规则
	// *act 模糊匹配规则 URI层级可以为空
	// {page} 字段匹配规则
	s2.BindHandler("GET:/{class}-{course}/:name/*act", func(r *ghttp.Request) {
		r.Response.Writef(
			"%v %v %v %v", r.Get("class"), r.Get("course"), r.Get("name"), r.Get("act"),
		)
	})
	// 绑定路由方法
	// s2.BindObjectMethod("/demo2", api.User, "Demo2")
	// RESTful对象注册
	// s2.BindObjectRest("/demo3", api.Demo3)
	s3 := g.Server("s3")
	s3.SetNameToUriType(ghttp.URI_TYPE_DEFAULT)
	s3.SetPort(8300)

	// 对象处理 请求校验
	type RegisterReq struct {
		Name  string `p:"username"  v:"required|length:6,30 #请输入账号 |账号长度为:min到:max位"`
		Pass  string `p:"password1" v:"required|length:6,30 #请输入密码 |密码长度不够"`                          // 增加v校验标签
		Pass2 string `p:"password2" v:"required|length:6,30| same:password1 #请确认密码 |密码长度不够 |两次密码不一致"` //tag p 标签来指定该属性绑定的参数名称
	}
	type RegisterRes struct {
		Code  int         `json:"code"`
		Error string      `json:"error"`
		Data  interface{} `json:"data"`
	}
	s3.BindHandler("/register", func(r *ghttp.Request) {
		// r.Response.WriteJson(r.Router)
		var req *RegisterReq
		if err := r.Parse(&req); err != nil {
			// 当请求校验错误时，所有校验失败的错误都返回了
			// 当产生错误时，我们可以将校验错误转换为*gvalid.Error对象，随后可以通过灵活的方法控制错误的返回
			if v, ok := err.(*gvalid.Error); ok {
				r.Response.WriteJsonExit(RegisterRes{
					Code:  1,
					Error: v.FirstString(),
				})
			}
			r.Response.WriteJsonExit(RegisterRes{
				Code:  1,
				Error: err.Error(),
				// Data:  nil,
			})
		} else {
			r.Response.WriteJsonExit(RegisterRes{
				Data: req,
			})
		}
	})
	s1.Start()
	s2.Start()
	s3.Start()
}
