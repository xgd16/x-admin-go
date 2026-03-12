package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CreateUser
type CreateUserReq struct {
	g.Meta   `path:"/user/create" tags:"User" method:"post" summary:"新增用户"`
	Username string `json:"username" v:"required|length:3,32#请输入用户名|用户名3-32字符"`
	Password string `json:"password" v:"required|length:6,32#请输入密码|密码6-32字符"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status" d:"1"`
}

type CreateUserRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id"`
}

// DeleteUser
type DeleteUserReq struct {
	g.Meta `path:"/user/delete" tags:"User" method:"post" summary:"删除用户"`
	Id     uint64 `json:"id" v:"required#请指定用户ID"`
}

type DeleteUserRes struct {
	g.Meta `mime:"application/json"`
}

// GetUser
type GetUserReq struct {
	g.Meta `path:"/user/get" tags:"User" method:"get" summary:"获取用户详情"`
	Id     uint64 `json:"id" v:"required#请指定用户ID"`
}

type GetUserRes struct {
	g.Meta   `mime:"application/json"`
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

// GetUserList
type GetUserListReq struct {
	g.Meta   `path:"/user/list" tags:"User" method:"post" summary:"用户列表"`
	PageNum  int    `json:"pageNum" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   *int   `json:"status"`
}

type GetUserListRes struct {
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

// UpdateUser
type UpdateUserReq struct {
	g.Meta   `path:"/user/update" tags:"User" method:"post" summary:"更新用户"`
	Id       uint64  `json:"id" v:"required#请指定用户ID"`
	Username string  `json:"username" v:"length:3,32#用户名3-32字符"`
	Password string  `json:"password" v:"length:6,32#密码6-32字符"`
	Nickname *string `json:"nickname"`
	Avatar   *string `json:"avatar"`
	Status   *int    `json:"status"`
}

type UpdateUserRes struct {
	g.Meta `mime:"application/json"`
}
