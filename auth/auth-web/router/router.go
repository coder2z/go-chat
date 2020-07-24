package router

import (
	authSvr "auth/auth-svr/proto/auth-svr"
	"auth/auth-web/handler"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

func InitRouter() *gin.Engine {

	oAuth := &handler.Auth{
		AuthSvr: authSvr.NewAuthSvrService("go.micro.chat.auth.service", client.DefaultClient),
	}

	//gin
	app := gin.Default()
	api := app.Group("/auth")
	{
		api.POST("/login", oAuth.Login)
		api.POST("/register", oAuth.Register)
		api.GET("/me", oAuth.Info)

	}
	return app
}
