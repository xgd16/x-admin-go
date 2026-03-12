package code

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
)

// 不需要记录到系统日志
var (
	UserExist        = newCode("用户已存在", &baseStartCode)
	UserNotExist     = newCode("账号或密码错误", &baseStartCode)
	UserNotFound     = newCode("用户不存在", &baseStartCode)
	UserDisabled     = newCode("用户已禁用", &baseStartCode)
	LoginFailMessage = newCode("账号或密码错误", &baseStartCode)
	InvalidToken     = newCode("token 无效或已过期", &baseStartCode)
	NotLoggedIn      = newCode("未登录", &baseStartCode)
	WrongOldPassword = newCode("原密码错误", &baseStartCode)
)

// 需要记录到日志
var (
	FailedToQueryUserDataProcedure = newCode("用户数据查询失败", &loggerStartCode)
	FailedToRequest                = newCode("请求失败", &loggerStartCode)
)

// CustomCode 自定义错误信息
func CustomCode(code gcode.Code, message string) gcode.Code {
	return gcode.New(code.Code(), fmt.Sprintf("%s - %s", code.Message(), message), code.Detail())
}
