package config

import (
	"common/config"
	"encoding/json"
)

type RabbitMq struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var (
	RabbitMqConfig *RabbitMq
)

func InitConfig() (err error) {
	configClient := config.GetConfigClient()
	//mysql
	rabbitMqConfig := configClient.GetNameSpace("rabbitMq.json")
	err = json.Unmarshal([]byte(rabbitMqConfig["content"].(string)), &RabbitMqConfig)
	return
}
