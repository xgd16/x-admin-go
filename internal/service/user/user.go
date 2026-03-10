package user

import (
	"context"

	"x-admin/api/user/v1"
)

type (
	IUser interface {
		GetList(ctx context.Context, req *v1.GetListReq) (*v1.GetListRes, error)
		Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error)
		Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error)
		Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error)
		Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error)
	}
)

var localUser IUser

func RegisterUser(u IUser) {
	localUser = u
}

func User() IUser {
	if localUser == nil {
		panic("user service 未注册")
	}
	return localUser
}
