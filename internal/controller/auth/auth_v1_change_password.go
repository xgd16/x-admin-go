package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"x-admin/api/auth/v1"
	"x-admin/internal/middleware"
	authService "x-admin/internal/service/auth"
)

func (c *ControllerV1) ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {
	user := middleware.GetUser(ghttp.RequestFromCtx(ctx))
	if user == nil {
		return nil, gerror.New("未登录")
	}
	if err := authService.Auth().ChangePassword(ctx, user, req.OldPassword, req.NewPassword); err != nil {
		return nil, err
	}
	return &v1.ChangePasswordRes{Message: "修改成功"}, nil
}
