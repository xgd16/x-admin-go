// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAdmin is the golang structure of table sys_admin for DAO operations like Where/Data.
type SysAdmin struct {
	g.Meta     `orm:"table:sys_admin, do:true"`
	Id         any         //
	Username   any         // 用户名
	Password   any         // 密码(bcrypt)
	Nickname   any         // 昵称
	Avatar     any         // 头像
	Status     any         // 状态: 0禁用 1启用
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
