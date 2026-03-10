// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAdmin is the golang structure for table sys_admin.
type SysAdmin struct {
	Id         uint64      `json:"id"         orm:"id"         description:""`            //
	Username   string      `json:"username"   orm:"username"   description:"用户名"`         // 用户名
	Password   string      `json:"password"   orm:"password"   description:"密码(bcrypt)"`  // 密码(bcrypt)
	Nickname   string      `json:"nickname"   orm:"nickname"   description:"昵称"`          // 昵称
	Avatar     string      `json:"avatar"     orm:"avatar"     description:"头像"`          // 头像
	Status     int         `json:"status"     orm:"status"     description:"状态: 0禁用 1启用"` // 状态: 0禁用 1启用
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" description:"创建时间"`        // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" description:"更新时间"`        // 更新时间
}
