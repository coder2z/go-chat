package main

import (
	"auth/auth-svr/config"
	"auth/auth-svr/handler"
	"auth/auth-svr/models"
	"auth/auth-svr/repositories"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"

	authsvr "auth/auth-svr/proto/auth-svr"
)

func main() {

	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	if err := models.InitModel(); err != nil {
		log.Fatal(err)
	}

	// New Service
	svr := grpc.NewService(
		service.Name("go.micro.chat.auth.service"),
		service.Version("latest"),
	)

	// Initialise service
	svr.Init()

	// Register Handler
	_ = authsvr.RegisterAuthSvrHandler(svr.Server(), &handler.AuthSvr{
		UserRepository: repositories.NewUserRepository(),
	})

	// Run service
	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
