package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateReq struct {
	g.Meta   `path:"/user/update" tags:"User" method:"post" summary:"更新用户"`
	Id       uint64 `json:"id" v:"required#请指定用户ID"`
	Username string  `json:"username" v:"length:3,32#用户名3-32字符"`
	Password string  `json:"password" v:"length:6,32#密码6-32字符"`
	Nickname *string `json:"nickname"`
	Avatar   *string `json:"avatar"`
	Status   *int   `json:"status"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
}
