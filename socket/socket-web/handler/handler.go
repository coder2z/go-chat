package handler

import (
	"common/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	socket "socket-svr/proto/socket-svr"
	websocketConn "socket/socket-web/websocket"
)

type Socket struct {
	SocketSvr socket.SocketSvrService
	RabbitMq  *RabbitMQ
}

var upGrader = websocket.Upgrader{
	//允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ConnAll = make(map[*websocketConn.Websocket]bool)

func (s *Socket) Ws(c *gin.Context) {
	userInfo, exists := c.Get("jwtUserInfo")
	if exists {
		user := userInfo.(jwt.UserInfo)
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		conn, err := websocketConn.InitConnection(ws)
		ConnAll[conn] = true
		defer func() {
			delete(ConnAll, conn)
			conn.Close()
		}()

		for {
			data, _ := conn.ReadMessage()
			res, _ := s.SocketSvr.Send(c, &socket.Request{
				Message: user.Name + "：" + string(data),
			})
			s.RabbitMq.PublishPub(res.Message)
		}
	}
}
