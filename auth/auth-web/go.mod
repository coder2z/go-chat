module auth/auth-web

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.9.1
	common v0.0.0
	auth/auth-svr v0.0.0
)

replace auth/auth-svr v0.0.0 => ../../auth/auth-svr
replace common v0.0.0 => ../../common