package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ChangePasswordReq struct {
	g.Meta     `path:"/auth/change-password" tags:"Auth" method:"post" summary:"修改密码"`
	OldPassword string `json:"oldPassword" v:"required#请输入原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,32#请输入新密码|密码6-32字符"`
}

type ChangePasswordRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message,omitempty"`
}
