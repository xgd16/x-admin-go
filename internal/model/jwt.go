package model

// JWTGenTokenInput 生成 Token 入参
type JWTGenTokenInput struct {
	Subject string `json:"subject"` // 对应 config 中 jwt.{subject}，如 admin
	Id      int64  `json:"id"`     // 用户 ID，写入 claims
}

// JWTGenTokenOutput 生成 Token 出参
type JWTGenTokenOutput struct {
	Token  string `json:"token"`  // JWT 字符串
	Expire int64  `json:"expire"` // 过期时间戳（毫秒）
}

// JWTVerifyTokenInput 校验 Token 入参
type JWTVerifyTokenInput struct {
	Token   string `json:"token"`   // JWT 字符串
	Subject string `json:"subject"` // 对应 config 中 jwt.{subject}，如 admin
}

// JWTVerifyTokenOutput 校验 Token 出参
type JWTVerifyTokenOutput struct {
	Expire int64  `json:"expire"` // 过期时间戳（秒）
	Id     int64  `json:"id"`     // 用户 ID
	JTI    string `json:"jti"`    // 令牌 ID
}
