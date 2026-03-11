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

func (s *sAuth) Login(ctx context.Context, username, password string) (*entity.UserInfo, string, int64, error) {
	var dbUser entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, username).Scan(&dbUser)
	if err != nil || dbUser.Id == 0 || dbUser.Status != 1 {
		return nil, "", 0, gerror.New("用户名或密码错误")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
		return nil, "", 0, gerror.New("用户名或密码错误")
	}
	user := &entity.UserInfo{
		Id:       dbUser.Id,
		Username: dbUser.Username,
		Nickname: dbUser.Nickname,
		Avatar:   dbUser.Avatar,
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
	var dbUser entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, user.Username).Scan(&dbUser)
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}
	if dbUser.Id == 0 {
		return gerror.New("用户不存在")
	}
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
