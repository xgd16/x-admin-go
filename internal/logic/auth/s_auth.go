package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"x-admin/internal/code"
	"x-admin/internal/consts"
	"x-admin/internal/dao"
	"x-admin/internal/model"
	"x-admin/internal/model/do"
	"x-admin/internal/model/entity"
	"x-admin/internal/service"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s *sAuth) Login(ctx context.Context, username, password string) (*entity.UserInfo, string, int64, error) {
	var dbUser entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, username).Scan(&dbUser)
	if err != nil || dbUser.Id == 0 || dbUser.Status != 1 {
		return nil, "", 0, code.ToError(code.LoginFailMessage)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
		return nil, "", 0, code.ToError(code.LoginFailMessage)
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
	out, err := service.Jwt().GenToken(ctx, &model.JWTGenTokenInput{
		Subject: "admin",
		Id:      int64(user.Id),
	})
	if err != nil {
		return nil, "", 0, code.ToErrorWrap(err, code.FailedToRequest, "生成 Token 失败")
	}
	// 返回给前端的 expire 为剩余有效秒数（config 中 expire 单位为小时）
	opt, _ := consts.GetJwtOptions("admin")
	expireSec := int64(168) * 3600 // 默认 7 天
	if opt != nil {
		expireSec = int64(opt.Expire) * 3600
	}
	return user, out.Token, expireSec, nil
}

func (s *sAuth) VerifyToken(ctx context.Context, tokenStr string) (*entity.UserInfo, error) {
	out, valid, err := service.Jwt().VerifyToken(ctx, &model.JWTVerifyTokenInput{
		Token:   tokenStr,
		Subject: "admin",
	})
	if err != nil || !valid {
		return nil, code.ToError(code.InvalidToken)
	}
	var dbUser entity.SysAdmin
	err = dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Id, out.Id).Scan(&dbUser)
	if err != nil || dbUser.Id == 0 || dbUser.Status != 1 {
		return nil, code.ToError(code.InvalidToken)
	}
	return &entity.UserInfo{
		Id:       dbUser.Id,
		Username: dbUser.Username,
		Nickname: dbUser.Nickname,
		Avatar:   dbUser.Avatar,
	}, nil
}

func (s *sAuth) ChangePassword(ctx context.Context, user *entity.UserInfo, oldPassword, newPassword string) error {
	var dbUser entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, user.Username).Scan(&dbUser)
	if err != nil {
		return code.ToErrorWrap(err, code.FailedToQueryUserDataProcedure, "查询用户失败")
	}
	if dbUser.Id == 0 {
		return code.ToError(code.UserNotFound)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(oldPassword)); err != nil {
		return code.ToError(code.WrongOldPassword)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return code.ToErrorWrap(err, code.FailedToRequest, "密码加密失败")
	}
	_, err = dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Id, dbUser.Id).Data(do.SysAdmin{Password: string(hash)}).Update()
	if err != nil {
		return code.ToErrorWrap(err, code.FailedToRequest, "更新密码失败")
	}
	return nil
}
