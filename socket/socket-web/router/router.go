package router

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	socketSvr "socket-svr/proto/socket-svr"
	"socket/socket-web/handler"
)

func InitRouter() *gin.Engine {
	socket := &handler.Socket{
		SocketSvr: socketSvr.NewSocketSvrService("go.micro.chat.socket.service", client.DefaultClient),
		RabbitMq:  handler.NewRabbitMQPubSub("chatMessage"),
	}
	//gin
	app := gin.Default()
	api := app.Group("/socket")
	{
		api.GET("/ws", socket.Ws)
	}
	return app
}
