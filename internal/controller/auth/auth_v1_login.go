package auth

import (
	"context"

	"x-admin/api/auth/v1"
	"x-admin/internal/service/auth"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	_, token, expire, err := auth.Auth().Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		Token:  token,
		Expire: expire,
		Message: "登录成功",
	}, nil
}
