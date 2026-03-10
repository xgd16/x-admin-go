package user

import (
	"context"

	"x-admin/api/user/v1"
	"x-admin/internal/service/user"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return user.User().Update(ctx, req)
}
