package entity

// UserInfo 当前登录用户信息（存入 Context）
type UserInfo struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
