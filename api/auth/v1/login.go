package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/auth/login" tags:"Auth" method:"post" summary:"后台登录"`
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}

type LoginRes struct {
	g.Meta  `mime:"application/json"`
	Token   string `json:"token"`
	Expire  int64  `json:"expire"`
	Message string `json:"message,omitempty"`
}
