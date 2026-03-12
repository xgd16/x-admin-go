// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"x-admin/internal/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// AuthJwt 返回 JWT 校验中间件，subject 对应 config 中 jwt.{subject}（如 admin）
		// 若 handler 的 Req 定义中 g.Meta 含 noAuth:"true"，则跳过校验
		AuthJwt(subject string) func(r *ghttp.Request)
		// GetUser 从 Request 上下文中获取当前用户（需在 AuthJwt 之后调用）
		GetUser(r *ghttp.Request) *entity.UserInfo
		// HandlerResponse 处理 Http 请求返回结果
		HandlerResponse(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
