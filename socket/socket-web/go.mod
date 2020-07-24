module socket/socket-web

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/websocket v1.4.1
	github.com/micro/go-micro/v2 v2.9.1
	github.com/streadway/amqp v1.0.0
	google.golang.org/protobuf v1.22.0
	socket-svr v0.0.0
	common v0.0.0
)

replace (
    socket-svr v0.0.0 => ../socket-svr
    common v0.0.0 => ../../common
)
