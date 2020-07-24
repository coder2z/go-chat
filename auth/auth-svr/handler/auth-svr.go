package handler

import (
	"auth/auth-svr/models"
	authsvr "auth/auth-svr/proto/auth-svr"
	"auth/auth-svr/repositories"
	"common/jwt"
	"context"
)

type AuthSvr struct {
	UserRepository repositories.UserRepositoryImp
}

func (s *AuthSvr) Login(c context.Context, in *authsvr.Request, out *authsvr.LoginResponse) (err error) {
	user, err := s.UserRepository.GetUserByName(in.Name)
	if err != nil {
		out.Msg = "用户名或者密码错误！"
		out.Code = 0
		return nil
	}
	ok := user.CheckPassword(in.Password)
	if !ok {
		out.Msg = "用户名或者密码错误！"
		out.Code = 0
		return nil
	}
	info := jwt.UserInfo{
		Id:   user.ID,
		Name: user.Name,
	}
	out.Token, err = info.GenerateToken()
	if err != nil {
		out.Msg = "服务器错误！"
		out.Code = 0
		return nil
	}
	out.Code = 1
	out.Msg = "success"
	return nil
}

func (s *AuthSvr) Info(c context.Context, in *authsvr.InfoRequest, out *authsvr.InfoResponse) (err error) {
	info := jwt.UserInfo{}
	err = info.ParseToken(in.Token)
	if err != nil {
		out.Msg = "请先登录"
		out.Code = 0
		return nil
	}
	out.Data = &authsvr.UserInfo{
		Name: info.Name,
		Id:   int64(info.Id),
	}
	out.Msg = "success"
	out.Code = 1
	return nil
}

func (s *AuthSvr) Register(c context.Context, in *authsvr.Request, out *authsvr.RegisterResponse) (err error) {
	user := &models.User{
		Name:     in.Name,
		Password: in.Password,
	}
	err = user.SetPassword()
	if err != nil {
		out.Msg = "服务器错误！"
		out.Code = 0
		return
	}
	err = s.UserRepository.AddUser(user)
	if err != nil {
		out.Msg = "服务器错误！"
		out.Code = 0
		return
	}
	out.Msg = "success"
	out.Code = 1
	return
}
