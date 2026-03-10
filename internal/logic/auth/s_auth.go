package auth

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"x-admin/internal/dao"
	"x-admin/internal/model/do"
	"x-admin/internal/model/entity"
	authService "x-admin/internal/service/auth"
)

type sAuth struct{}

func init() {
	authService.RegisterAuth(New())
}

func New() authService.IAuth {
	return &sAuth{}
}

// 预生成的 "123456" 的 bcrypt 哈希 (cost=10)
const defaultPasswordHash = "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"

func (s *sAuth) Login(ctx context.Context, username, password string) (*entity.UserInfo, string, int64, error) {
	// 1. 优先检查 sys_admin（含通过修改密码迁移的配置 admin）
	var dbUser entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, username).Scan(&dbUser)
	if err == nil && dbUser.Id > 0 && dbUser.Status == 1 {
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err == nil {
			user := &entity.UserInfo{
				Id:       dbUser.Id,
				Username: dbUser.Username,
				Nickname: dbUser.Nickname,
				Avatar:   dbUser.Avatar,
			}
			return s.issueToken(ctx, user)
		}
	}

	// 2. 回退到配置认证
	cfg := g.Cfg().MustGet(ctx, "auth")
	adminUser := cfg.Map()["admin"]
	if adminUser == nil {
		return nil, "", 0, gerror.New("auth.admin 未配置")
	}
	adminMap, ok := adminUser.(map[string]interface{})
	if !ok || len(adminMap) == 0 {
		return nil, "", 0, gerror.New("auth.admin 配置格式错误")
	}
	cfgUsername, _ := adminMap["username"].(string)
	if cfgUsername == "" {
		cfgUsername = "admin"
	}
	cfgPassword := ""
	if v, ok := adminMap["password"].(string); ok {
		cfgPassword = v
	}
	if username != cfgUsername {
		return nil, "", 0, gerror.New("用户名或密码错误")
	}
	cfgPasswordHash, hasHash := adminMap["passwordHash"].(string)
	if hasHash && cfgPasswordHash != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(cfgPasswordHash), []byte(password)); err != nil {
			return nil, "", 0, gerror.New("用户名或密码错误")
		}
	} else if cfgPassword != "" {
		if password != cfgPassword {
			return nil, "", 0, gerror.New("用户名或密码错误")
		}
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(defaultPasswordHash), []byte(password)); err != nil {
			return nil, "", 0, gerror.New("用户名或密码错误")
		}
	}
	user := &entity.UserInfo{
		Id:       1,
		Username: username,
		Nickname: username,
		Avatar:   "",
	}
	return s.issueToken(ctx, user)
}

func (s *sAuth) issueToken(ctx context.Context, user *entity.UserInfo) (*entity.UserInfo, string, int64, error) {
	secret := g.Cfg().MustGet(ctx, "auth.jwtSecret").String()
	if secret == "" {
		secret = "x-admin-jwt-secret-change-in-production"
	}
	expire := g.Cfg().MustGet(ctx, "auth.jwtExpire").Int64()
	if expire <= 0 {
		expire = 86400 * 7 // 7 天
	}
	claims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"nickname": user.Nickname,
		"exp":      time.Now().Add(time.Duration(expire) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, "", 0, gerror.Wrap(err, "生成 Token 失败")
	}
	return user, tokenStr, expire, nil
}

func (s *sAuth) VerifyToken(ctx context.Context, tokenStr string) (*entity.UserInfo, error) {
	secret := g.Cfg().MustGet(ctx, "auth.jwtSecret").String()
	if secret == "" {
		secret = "x-admin-jwt-secret-change-in-production"
	}
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, gerror.New("token 无效或已过期")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, gerror.New("token 无效")
	}
	id, _ := claims["id"].(float64)
	username, _ := claims["username"].(string)
	nickname, _ := claims["nickname"].(string)
	return &entity.UserInfo{
		Id:       uint64(id),
		Username: username,
		Nickname: nickname,
	}, nil
}

func (s *sAuth) ChangePassword(ctx context.Context, user *entity.UserInfo, oldPassword, newPassword string) error {
	// 1. 检查 sys_admin 中是否存在该用户
	var dbUser entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, user.Username).Scan(&dbUser)
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}

	if dbUser.Id > 0 {
		// 已在 sys_admin：验证原密码并更新
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(oldPassword)); err != nil {
			return gerror.New("原密码错误")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return gerror.Wrap(err, "密码加密失败")
		}
		_, err = dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Id, dbUser.Id).Data(do.SysAdmin{Password: string(hash)}).Update()
		if err != nil {
			return gerror.Wrap(err, "更新密码失败")
		}
		return nil
	}

	// 2. 配置 admin：验证原密码后迁移到 sys_admin
	cfg := g.Cfg().MustGet(ctx, "auth")
	adminMap, _ := cfg.Map()["admin"].(map[string]interface{})
	if adminMap == nil {
		return gerror.New("您的账户不支持修改密码")
	}
	cfgUsername, _ := adminMap["username"].(string)
	if cfgUsername == "" {
		cfgUsername = "admin"
	}
	if user.Username != cfgUsername {
		return gerror.New("您的账户不支持修改密码")
	}
	cfgPassword, _ := adminMap["password"].(string)
	cfgPasswordHash, hasHash := adminMap["passwordHash"].(string)
	oldOk := false
	if hasHash && cfgPasswordHash != "" {
		oldOk = bcrypt.CompareHashAndPassword([]byte(cfgPasswordHash), []byte(oldPassword)) == nil
	} else if cfgPassword != "" {
		oldOk = oldPassword == cfgPassword
	} else {
		oldOk = bcrypt.CompareHashAndPassword([]byte(defaultPasswordHash), []byte(oldPassword)) == nil
	}
	if !oldOk {
		return gerror.New("原密码错误")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return gerror.Wrap(err, "密码加密失败")
	}
	_, err = dao.SysAdmin.Ctx(ctx).Data(do.SysAdmin{
		Username: user.Username,
		Password: string(hash),
		Nickname: user.Nickname,
		Status:   1,
	}).Insert()
	if err != nil {
		return gerror.Wrap(err, "更新密码失败")
	}
	return nil
}
