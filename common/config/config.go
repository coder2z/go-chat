package config

import "github.com/shima-park/agollo"

const Apollo = "docker.myxy99.cn:8080"

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
