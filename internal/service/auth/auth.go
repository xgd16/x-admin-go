package auth

import (
	"context"

	"x-admin/internal/model/entity"
)

type (
	IAuth interface {
		Login(ctx context.Context, username, password string) (*entity.UserInfo, string, int64, error)
		VerifyToken(ctx context.Context, token string) (*entity.UserInfo, error)
		ChangePassword(ctx context.Context, user *entity.UserInfo, oldPassword, newPassword string) error
	}
)

var localAuth IAuth

func RegisterAuth(a IAuth) {
	localAuth = a
}

func Auth() IAuth {
	if localAuth == nil {
		panic("auth service 未注册")
	}
	return localAuth
}
