package global

import (
	"go-temp/conf"
	"gorm.io/gorm"
)

var (
	AppConfig    *conf.AppsConfig
	DB           *gorm.DB
	ENV          *string
	ServerConfig *conf.ServerConfig
)
