package user

import (
	"context"

	"x-admin/api/user/v1"
	"x-admin/internal/service/user"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return user.User().Delete(ctx, req)
}
