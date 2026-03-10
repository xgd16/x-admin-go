package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"x-admin/api/auth/v1"
	"x-admin/internal/middleware"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	user := middleware.GetUser(ghttp.RequestFromCtx(ctx))
	if user == nil {
		return nil, gerror.New("未登录")
	}
	return &v1.GetInfoRes{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}
