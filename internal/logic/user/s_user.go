package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	v1 "x-admin/api/admin/v1"
	"x-admin/internal/code"
	"x-admin/internal/dao"
	"x-admin/internal/model/do"
	"x-admin/internal/model/entity"
	"x-admin/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) GetUserList(ctx context.Context, req *v1.GetUserListReq) (*v1.GetUserListRes, error) {
	m := dao.SysAdmin.Ctx(ctx)
	if req.Username != "" {
		m = m.WhereLike(dao.SysAdmin.Columns().Username, "%"+req.Username+"%")
	}
	if req.Nickname != "" {
		m = m.WhereLike(dao.SysAdmin.Columns().Nickname, "%"+req.Nickname+"%")
	}
	if req.Status != nil {
		m = m.Where(dao.SysAdmin.Columns().Status, *req.Status)
	}
	total, err := m.Clone().Count()
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToQueryUserDataProcedure, "统计用户数失败")
	}
	pageNum, pageSize := req.PageNum, req.PageSize
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	var list []*entity.SysAdmin
	err = m.Page(pageNum, pageSize).OrderDesc(dao.SysAdmin.Columns().Id).Scan(&list)
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToQueryUserDataProcedure, "查询用户列表失败")
	}
	items := make([]v1.UserItem, 0, len(list))
	for _, e := range list {
		createTime := ""
		if e.CreateTime != nil {
			createTime = e.CreateTime.Format("Y-m-d H:i:s")
		}
		items = append(items, v1.UserItem{
			Id:         e.Id,
			Username:   e.Username,
			Nickname:   e.Nickname,
			Avatar:     e.Avatar,
			Status:     e.Status,
			CreateTime: createTime,
		})
	}
	return &v1.GetUserListRes{List: items, Total: int64(total)}, nil
}

func (s *sUser) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserRes, error) {
	var e entity.SysAdmin
	err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Id, req.Id).Scan(&e)
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToQueryUserDataProcedure, "查询用户失败")
	}
	if e.Id == 0 {
		return nil, code.ToError(code.UserNotFound)
	}
	return &v1.GetUserRes{
		Id:       e.Id,
		Username: e.Username,
		Nickname: e.Nickname,
		Avatar:   e.Avatar,
		Status:   e.Status,
	}, nil
}

func (s *sUser) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserRes, error) {
	n, err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, req.Username).Count()
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToQueryUserDataProcedure, "检查用户名失败")
	}
	if n > 0 {
		return nil, code.ToError(code.UserExist)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToRequest, "密码加密失败")
	}
	data := do.SysAdmin{
		Username: req.Username,
		Password: string(hashed),
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   req.Status,
	}
	if data.Status == 0 {
		data.Status = 1
	}
	id, err := dao.SysAdmin.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToRequest, "创建用户失败")
	}
	return &v1.CreateUserRes{Id: uint64(id)}, nil
}

func (s *sUser) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserRes, error) {
	data := make(map[string]interface{})
	if req.Username != "" {
		n, _ := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Username, req.Username).WhereNot(dao.SysAdmin.Columns().Id, req.Id).Count()
		if n > 0 {
			return nil, code.ToError(code.UserExist)
		}
		data[dao.SysAdmin.Columns().Username] = req.Username
	}
	if req.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, code.ToErrorWrap(err, code.FailedToRequest, "密码加密失败")
		}
		data[dao.SysAdmin.Columns().Password] = string(hashed)
	}
	if req.Nickname != nil {
		data[dao.SysAdmin.Columns().Nickname] = *req.Nickname
	}
	if req.Avatar != nil {
		data[dao.SysAdmin.Columns().Avatar] = *req.Avatar
	}
	if req.Status != nil {
		data[dao.SysAdmin.Columns().Status] = *req.Status
	}
	if len(data) == 0 {
		return &v1.UpdateUserRes{}, nil
	}
	_, err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToRequest, "更新用户失败")
	}
	return &v1.UpdateUserRes{}, nil
}

func (s *sUser) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserRes, error) {
	_, err := dao.SysAdmin.Ctx(ctx).Where(dao.SysAdmin.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, code.ToErrorWrap(err, code.FailedToRequest, "删除用户失败")
	}
	return &v1.DeleteUserRes{}, nil
}
