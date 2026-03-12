package admin

import (
	"context"

	v1 "x-admin/api/admin/v1"
	"x-admin/internal/service"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	return service.User().CreateUser(ctx, req)
}

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	return service.User().DeleteUser(ctx, req)
}

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	return service.User().GetUser(ctx, req)
}

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	return service.User().GetUserList(ctx, req)
}

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	return service.User().UpdateUser(ctx, req)
}
