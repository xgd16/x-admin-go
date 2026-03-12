package model

// JwtOption JWT 配置项，对应 config.yaml 中 jwt.{subject} 结构
type JwtOption struct {
	Secret  string `json:"secret"`  // 签名密钥
	Expire  int    `json:"expire"`  // 过期时间，单位：小时
	Issuer  string `json:"issuer"`  // 签发者
	Subject string `json:"subject"` // 主题
}

// ConfigData 系统配置结构，由 consts 加载
type ConfigData struct {
	Jwt map[string]JwtOption `json:"jwt"` // key 为 subject，如 admin
}
