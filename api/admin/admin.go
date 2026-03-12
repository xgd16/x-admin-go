// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"x-admin/api/admin/v1"
)

type IAdminV1 interface {
	LoginAuth(ctx context.Context, req *v1.LoginAuthReq) (res *v1.LoginAuthRes, err error)
	GetAuthInfo(ctx context.Context, req *v1.GetAuthInfoReq) (res *v1.GetAuthInfoRes, err error)
	ChangePasswordAuth(ctx context.Context, req *v1.ChangePasswordAuthReq) (res *v1.ChangePasswordAuthRes, err error)
	CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
	GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
}
