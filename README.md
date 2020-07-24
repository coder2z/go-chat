# 基于go-micro开发的微服务聊天室

## 技术栈
> 微服务框架：go-micro

> web框架：gin

> orm:gorm

> websocket

> 配置中心：apollo

> jwt


## 网站架构

![网站架构](https://gitee.com/myxy99/pic/raw/master/img/blog/2020/07/24/20200724132536.png)

## 项目启动

修改common/config/config.go 中的apollo的host

apollo中添加配置

![](https://gitee.com/myxy99/pic/raw/master/img/blog/2020/07/24/20200724140425.png)

启动认证服务
```
cd auth/auth-svr
go run main.go
```

启动认证服务api
```
cd auth/auth-web
go run main.go
```

启动socket服务
```
cd socket/socket-svr
go run main.go
```

启动socket服务api
```
cd socket/socket-web
go run main.go
```