package handler

import (
	"context"
	. "socket/socket-svr/proto/socket-svr"
)

type SocketSvr struct{}

func (s *SocketSvr) Send(ctx context.Context, in *Request, out *Response) (err error) {
	//简单处理，这里就直接赋值，不处理了
	out.Message = in.Message
	return
}
