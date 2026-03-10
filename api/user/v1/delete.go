package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteReq struct {
	g.Meta `path:"/user/delete" tags:"User" method:"post" summary:"删除用户"`
	Id     uint64 `json:"id" v:"required#请指定用户ID"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
}
