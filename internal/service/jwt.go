// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"x-admin/internal/model"
)

type (
	IJwt interface {
		// GenToken 生成 token，使用 config.yaml 中 jwt.{subject} 配置
		GenToken(ctx context.Context, in *model.JWTGenTokenInput) (out *model.JWTGenTokenOutput, err error)
		VerifyToken(ctx context.Context, in *model.JWTVerifyTokenInput) (out *model.JWTVerifyTokenOutput, valid bool, err error)
	}
)

var (
	localJwt IJwt
)

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}
