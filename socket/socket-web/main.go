package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"socket/socket-web/config"
	"socket/socket-web/handler"
	"socket/socket-web/router"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	// create new web service
	service := web.NewService(
		web.Name("go.micro.api.socket"),
		web.Version("latest"),
	)

	go func() {
		rabbitmq := handler.NewRabbitMQPubSub("chat")
		rabbitmq.RecieveSub()
	}()

	// initialise service
	_ = service.Init()

	app := router.InitRouter()

	// register html handler
	service.Handle("/", app)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
