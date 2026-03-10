package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta   `path:"/user/list" tags:"User" method:"post" summary:"用户列表"`
	PageNum  int    `json:"pageNum" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   *int   `json:"status"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	List   []UserItem `json:"list"`
	Total  int64      `json:"total"`
}

type UserItem struct {
	Id         uint64 `json:"id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime,omitempty"`
}
