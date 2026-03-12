// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "x-admin/api/admin/v1"
)

type (
	IUser interface {
		GetUserList(ctx context.Context, req *v1.GetUserListReq) (*v1.GetUserListRes, error)
		GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserRes, error)
		CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserRes, error)
		UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserRes, error)
		DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserRes, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
