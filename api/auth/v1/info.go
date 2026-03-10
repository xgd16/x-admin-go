package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta `path:"/auth/info" tags:"Auth" method:"get" summary:"获取当前登录用户信息"`
}

type GetInfoRes struct {
	g.Meta  `mime:"application/json"`
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar,omitempty"`
}
