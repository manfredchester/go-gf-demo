package service

import (
	"github.com/gogf/gf/net/ghttp"
)

type serivceMiddleware struct {
}

var Middleware = new(serivceMiddleware)

func (e *serivceMiddleware) Ctx(r *ghttp.Request) {

}
func (e *serivceMiddleware) Auth(r *ghttp.Request) {

}
