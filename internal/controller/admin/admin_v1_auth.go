package admin

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	v1 "x-admin/api/admin/v1"
	"x-admin/internal/code"
	"x-admin/internal/middleware"
	"x-admin/internal/service"
)

func (c *ControllerV1) LoginAuth(ctx context.Context, req *v1.LoginAuthReq) (res *v1.LoginAuthRes, err error) {
	_, token, expire, err := service.Auth().Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginAuthRes{
		Token:   token,
		Expire:  expire,
		Message: "登录成功",
	}, nil
}

func (c *ControllerV1) GetAuthInfo(ctx context.Context, req *v1.GetAuthInfoReq) (res *v1.GetAuthInfoRes, err error) {
	user := middleware.GetUser(ghttp.RequestFromCtx(ctx))
	if user == nil {
		return nil, code.ToError(code.NotLoggedIn)
	}
	return &v1.GetAuthInfoRes{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}

func (c *ControllerV1) ChangePasswordAuth(ctx context.Context, req *v1.ChangePasswordAuthReq) (res *v1.ChangePasswordAuthRes, err error) {
	user := middleware.GetUser(ghttp.RequestFromCtx(ctx))
	if user == nil {
		return nil, code.ToError(code.NotLoggedIn)
	}
	if err := service.Auth().ChangePassword(ctx, user, req.OldPassword, req.NewPassword); err != nil {
		return nil, err
	}
	return &v1.ChangePasswordAuthRes{Message: "修改成功"}, nil
}
