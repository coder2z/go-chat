package config

import "github.com/shima-park/agollo"

// Apollo host
const Apollo = "xxxx:8080"

func initConfig() agollo.Agollo {
	a, err := agollo.New(Apollo, "chat", agollo.AutoFetchOnCacheMiss())
	if err != nil {
		panic(err)
	}
	return a
}

//		a.Get("rabbitmq_url", agollo.WithNamespace("common")),

func GetConfigClient() agollo.Agollo {
	return initConfig()
}
