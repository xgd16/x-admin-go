package consts

import (
	"errors"

	"x-admin/internal/model"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Ctx          = gctx.New()
	Logger       = g.Log()
	SystemConfig *gjson.Json
	Config       *model.ConfigData
)

func init() {
	if err := initSystemConfig(); err != nil {
		panic(err)
	}
}

func initSystemConfig() (err error) {
	data, err := g.Cfg().Data(Ctx)
	if err != nil {
		return
	}
	SystemConfig = gjson.New(data)
	if err = SystemConfig.Scan(&Config); err != nil {
		return
	}
	return
}

// GetJwtOptions 根据 subject（如 admin）获取 JWT 配置，对应 config.yaml 中 jwt.{subject}
func GetJwtOptions(subject string) (*model.JwtOption, error) {
	if Config == nil || Config.Jwt == nil {
		return nil, errors.New("jwt config not loaded")
	}
	opt, ok := Config.Jwt[subject]
	if !ok {
		return nil, errors.New("jwt config not found for subject: " + subject)
	}
	return &opt, nil
}
