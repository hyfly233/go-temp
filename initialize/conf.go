package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppsConfig struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
	} `mapstructure:"server"`
}

var AppConfig *AppsConfig

func InitConf(env *string) {
	zap.S().Info("初始化配置文件 ------------------")
	v := viper.New()
	v.SetConfigFile(fmt.Sprintf("application-%s.yaml", *env))

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	AppConfig = new(AppsConfig)
	if err := v.Unmarshal(AppConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置文件初始化成功 ------------------")
}
