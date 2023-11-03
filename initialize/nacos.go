package initialize

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"

	"go-temp/conf"
	"go-temp/global"
)

func InitNacos(_ *string) {
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

	global.NacosConfig = new(conf.NacosConfig)
	err = json.Unmarshal([]byte(config), &global.NacosConfig)
	if err != nil {
		panic(err)
	}

	zap.S().Info("初始化配置中心成功 ------------------")
}
