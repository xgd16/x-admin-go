package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LoginAuth
type LoginAuthReq struct {
	g.Meta   `path:"/auth/login" tags:"Auth" method:"post" summary:"后台登录" noAuth:"true"`
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}

type LoginAuthRes struct {
	g.Meta  `mime:"application/json"`
	Token   string `json:"token"`
	Expire  int64  `json:"expire"`
	Message string `json:"message,omitempty"`
}

// GetAuthInfo
type GetAuthInfoReq struct {
	g.Meta `path:"/auth/info" tags:"Auth" method:"get" summary:"获取当前登录用户信息"`
}

type GetAuthInfoRes struct {
	g.Meta  `mime:"application/json"`
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar,omitempty"`
}

// ChangePasswordAuth
type ChangePasswordAuthReq struct {
	g.Meta      `path:"/auth/change-password" tags:"Auth" method:"post" summary:"修改密码"`
	OldPassword string `json:"oldPassword" v:"required#请输入原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,32#请输入新密码|密码6-32字符"`
}

type ChangePasswordAuthRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message,omitempty"`
}
