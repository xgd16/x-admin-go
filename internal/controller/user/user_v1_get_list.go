package user

import (
	"context"

	"x-admin/api/user/v1"
	"x-admin/internal/service/user"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return user.User().GetList(ctx, req)
}
