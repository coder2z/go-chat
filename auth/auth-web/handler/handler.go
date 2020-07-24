package handler

import (
	authSvr "auth/auth-svr/proto/auth-svr"
	"auth/auth-web/dao"
	R "common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	AuthSvr authSvr.AuthSvrService
}

func (a *Auth) Login(ctx *gin.Context) {
	var login dao.LoginDao
	if err := ctx.ShouldBind(&login); err == nil {
		req := authSvr.Request{
			Name:     login.Name,
			Password: login.Password,
		}
		res, err := a.AuthSvr.Login(ctx, &req)
		if err != nil {
			R.Response(ctx, http.StatusInternalServerError, "服务器错误", err.Error(), http.StatusInternalServerError)
			return
		}
		R.Response(ctx, int(res.Code), res.Msg, res.Token, http.StatusOK)
		return
	} else {
		R.Response(ctx, http.StatusUnprocessableEntity, "参数错误", err.Error(), http.StatusUnprocessableEntity)
		return
	}
}

func (a *Auth) Register(ctx *gin.Context) {
	var register dao.RegisterDao
	if err := ctx.ShouldBind(&register); err == nil {
		req := authSvr.Request{
			Name:     register.Name,
			Password: register.Password,
		}
		res, err := a.AuthSvr.Register(ctx, &req)
		if err != nil {
			R.Response(ctx, http.StatusInternalServerError, "服务器错误", err.Error(), http.StatusInternalServerError)
			return
		}
		R.Response(ctx, int(res.Code), res.Msg, nil, http.StatusOK)
		return
	} else {
		R.Response(ctx, http.StatusUnprocessableEntity, "参数错误", err.Error(), http.StatusUnprocessableEntity)
		return
	}
}

func (a *Auth) Info(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) < 7 {
		R.Response(ctx, http.StatusUnauthorized, "未登录", nil, http.StatusUnauthorized)
		return
	}
	req := authSvr.InfoRequest{
		Token: token[7:],
	}
	res, err := a.AuthSvr.Info(ctx, &req)
	if err != nil {
		R.Response(ctx, http.StatusInternalServerError, "服务器错误", err.Error(), http.StatusInternalServerError)
		return
	}
	if res.Code == 0 {
		R.Response(ctx, http.StatusUnauthorized, "未登录", nil, http.StatusUnauthorized)
		return
	} else {
		R.Response(ctx, int(res.Code), res.Msg, res.Data, http.StatusOK)
		return
	}
}
