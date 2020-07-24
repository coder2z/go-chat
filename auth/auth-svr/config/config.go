package config

import (
	"common/config"
	"encoding/json"
)

type Mysql struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var (
	MysqlConfig *Mysql
)

func InitConfig() (err error) {
	configClient := config.GetConfigClient()
	//mysql
	mysqlConfig := configClient.GetNameSpace("mysql.json")
	err = json.Unmarshal([]byte(mysqlConfig["content"].(string)), &MysqlConfig)


	return
}
