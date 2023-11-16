package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"go-temp/conf"
	"go-temp/global"
)

func InitConf(env *string) {
	zap.S().Info("获取 Nacos 配置文件 ------------------")
	v := viper.New()
	v.SetConfigFile(fmt.Sprintf("config-%s.yaml", *env))

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	global.NacosConfig = new(conf.NacosConfig)
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Info("获取 Nacos 配置文件成功 ------------------")

	zap.S().Info("初始化 Nacos 配置中心 ------------------")

	nsc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	ncc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
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
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}

	global.ServerConfig = new(conf.ServerConfig)
	err = json.Unmarshal([]byte(config), &global.ServerConfig)
	if err != nil {
		panic(err)
	}

	zap.S().Info("初始化 Nacos 配置中心成功 ------------------")
}
