// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAdminDao is the data access object for the table sys_admin.
type SysAdminDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysAdminColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysAdminColumns defines and stores column names for the table sys_admin.
type SysAdminColumns struct {
	Id         string //
	Username   string // 用户名
	Password   string // 密码(bcrypt)
	Nickname   string // 昵称
	Avatar     string // 头像
	Status     string // 状态: 0禁用 1启用
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// sysAdminColumns holds the columns for the table sys_admin.
var sysAdminColumns = SysAdminColumns{
	Id:         "id",
	Username:   "username",
	Password:   "password",
	Nickname:   "nickname",
	Avatar:     "avatar",
	Status:     "status",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
}

// NewSysAdminDao creates and returns a new DAO object for table data access.
func NewSysAdminDao(handlers ...gdb.ModelHandler) *SysAdminDao {
	return &SysAdminDao{
		group:    "default",
		table:    "sys_admin",
		columns:  sysAdminColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAdminDao) Columns() SysAdminColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAdminDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
