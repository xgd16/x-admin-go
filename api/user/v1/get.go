package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/user/get" tags:"User" method:"get" summary:"获取用户详情"`
	Id     uint64 `json:"id" v:"required#请指定用户ID"`
}

type GetRes struct {
	g.Meta   `mime:"application/json"`
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}
