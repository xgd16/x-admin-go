package user

import (
	"context"

	"x-admin/api/user/v1"
	"x-admin/internal/service/user"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return user.User().Create(ctx, req)
}
