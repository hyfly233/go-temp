package global

import (
	"gorm.io/gorm"

	"go-temp/conf"
)

var (
	AppConfig   *conf.AppsConfig
	DB          *gorm.DB
	ENV         *string
	NacosConfig *conf.NacosConfig
	GrpcConfig  *conf.GrpcConfig
)
