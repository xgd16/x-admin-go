package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"x-admin/internal/consts"
	"x-admin/internal/model/entity"
	"x-admin/internal/service/auth"
)

// AuthJwt JWT 校验中间件，将用户信息写入 Context
func AuthJwt(r *ghttp.Request) {
	token := r.GetHeader("Authorization")
	if token != "" && len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	if token == "" {
		token = r.Get("token").String()
	}
	if token == "" {
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{
			Code:    401,
			Message: "未登录或 token 已过期",
		})
		return
	}
	user, err := auth.Auth().VerifyToken(r.GetCtx(), token)
	if err != nil {
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{
			Code:    401,
			Message: gconv.String(err),
		})
		return
	}
	r.SetCtxVar(consts.ContextKeyUser, user)
	r.Middleware.Next()
}

// GetUser 从 Request 上下文中获取当前用户（需在 AuthJwt 之后调用）
func GetUser(r *ghttp.Request) *entity.UserInfo {
	v := r.GetCtxVar(consts.ContextKeyUser)
	if v.IsNil() {
		return nil
	}
	user, ok := v.Val().(*entity.UserInfo)
	if !ok {
		return nil
	}
	return user
}
