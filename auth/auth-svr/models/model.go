package models

import (
	"auth/auth-svr/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitModel() (err error) {
	DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.MysqlConfig.User,
			config.MysqlConfig.Password,
			config.MysqlConfig.Host,
			"chat"))

	if err != nil {
		return err
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	//注册模型

	DB.AutoMigrate(&User{})

	return DB.DB().Ping()
}
