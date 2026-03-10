package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"/user/create" tags:"User" method:"post" summary:"新增用户"`
	Username string `json:"username" v:"required|length:3,32#请输入用户名|用户名3-32字符"`
	Password string `json:"password" v:"required|length:6,32#请输入密码|密码6-32字符"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status" d:"1"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id"`
}
