package global

import (
	"gorm.io/gorm"

	"go-temp/conf"
)

var (
	NacosConfig  *conf.NacosConfig
	DB           *gorm.DB
	ENV          *string
	ServerConfig *conf.ServerConfig
)
