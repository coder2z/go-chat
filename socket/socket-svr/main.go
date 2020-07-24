package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"socket/socket-svr/handler"
	socketsvr "socket/socket-svr/proto/socket-svr"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.chat.socket.service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = socketsvr.RegisterSocketSvrHandler(service.Server(), new(handler.SocketSvr))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
