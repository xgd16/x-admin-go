// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"x-admin/internal/model/entity"
)

type (
	IAuth interface {
		Login(ctx context.Context, username string, password string) (*entity.UserInfo, string, int64, error)
		VerifyToken(ctx context.Context, tokenStr string) (*entity.UserInfo, error)
		ChangePassword(ctx context.Context, user *entity.UserInfo, oldPassword string, newPassword string) error
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
