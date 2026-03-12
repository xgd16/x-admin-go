package middleware

import (
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"x-admin/internal/code"
	"x-admin/internal/consts"
	"x-admin/internal/model/entity"
	"x-admin/internal/service"
)

// AuthJwt 返回 JWT 校验中间件，subject 对应 config 中 jwt.{subject}（如 admin）
// 若 handler 的 Req 定义中 g.Meta 含 noAuth:"true"，则跳过校验
func (s *sMiddleware) AuthJwt(subject string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		handler := r.GetServeHandler()
		if handler != nil && handler.GetMetaTag("noAuth") == "true" {
			r.Middleware.Next()
			return
		}

		authorization := r.GetHeader("Authorization")
		token := strings.TrimPrefix(authorization, "Bearer ")
		if token == "" {
			token = r.Get("token").String()
		}
		if token == "" {
			r.Response.Status = http.StatusUnauthorized
			r.Response.WriteJsonExit(g.Map{
				"code":        code.InvalidToken.Code(),
				"message":     code.InvalidToken.Message(),
				"serviceTime": gtime.Now().UnixMilli(),
			})
			return
		}

		user, err := service.Auth().VerifyToken(r.GetCtx(), token)
		if err != nil {
			r.Response.Status = http.StatusUnauthorized
			r.Response.WriteJsonExit(g.Map{
				"code":        code.InvalidToken.Code(),
				"message":     code.InvalidToken.Message(),
				"serviceTime": gtime.Now().UnixMilli(),
			})
			return
		}
		r.SetCtxVar(consts.ContextKeyUser, user)
		r.Middleware.Next()
	}
}

// GetUser 从 Request 上下文中获取当前用户（需在 AuthJwt 之后调用）
func (s *sMiddleware) GetUser(r *ghttp.Request) *entity.UserInfo {
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
