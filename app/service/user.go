package service

import "context"

// 中间件管理服务
var User = new(serviceUser)

type serviceUser struct{}

// 获得用户信息详情
func (s *serviceUser) GetProfile(ctx context.Context) {
	// *model.User {
	// return Session.GetUser(ctx)
}
