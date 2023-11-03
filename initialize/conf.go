package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"go-temp/conf"
	"go-temp/global"
)

func InitConf(env *string) {
	zap.S().Info("初始化配置文件 ------------------")
	v := viper.New()
	v.SetConfigFile(fmt.Sprintf("application-%s.yaml", *env))

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	global.AppConfig = new(conf.AppsConfig)
	if err := v.Unmarshal(global.AppConfig); err != nil {
		panic(err)
	}
	zap.S().Info("初始化配置文件成功 ------------------")

	global.NacosConfig = &global.AppConfig.NacosConfig
	global.GrpcConfig = &global.AppConfig.GrpcConfig

}
