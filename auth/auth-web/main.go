package main

import (
	"auth/auth-web/router"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.api.auth"),
		web.Version("latest"),
	)

	app := router.InitRouter()

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))
	service.Handle("/", app)
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
