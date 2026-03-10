package user

import (
	"context"

	"x-admin/api/user/v1"
	"x-admin/internal/service/user"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	return user.User().Get(ctx, req)
}
