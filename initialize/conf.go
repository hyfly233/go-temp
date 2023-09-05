package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go-temp/conf"
	"go-temp/global"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	zap.S().Info("初始化配置中心 ------------------")

	nsc := []constant.ServerConfig{
		{
			IpAddr: "192.168.5.24",
			Port:   8848,
		},
	}

	ncc := constant.ClientConfig{
		NamespaceId:         "123",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "logs/nacos/log",
		CacheDir:            "logs/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": nsc,
		"clientConfig":  ncc,
	})
	if err != nil {
		panic(err)
	}

	var config string
	config, err = client.GetConfig(vo.ConfigParam{
		DataId: "test123",
		Group:  "test",
	})
	if err != nil {
		panic(err)
	}

	global.ServerConfig = new(conf.ServerConfig)
	err = json.Unmarshal([]byte(config), &global.ServerConfig)
	if err != nil {
		panic(err)
	}

	zap.S().Info("初始化配置中心成功 ------------------")
}
