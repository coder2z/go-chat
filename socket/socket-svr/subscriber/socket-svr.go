package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	socketsvr "socket/socket-svr/proto/socket-svr"
)

type SocketSvr struct{}

func (e *SocketSvr) Handle(ctx context.Context, msg *socketsvr.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}
