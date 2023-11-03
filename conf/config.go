package conf

type (
	AppsConfig struct {
		Server      Server      `mapstructure:"server"`
		NacosConfig NacosConfig `mapstructure:"nacos"`
		GrpcConfig  GrpcConfig  `mapstructure:"grpc"`
	}

	Server struct {
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
	}

	NacosConfig struct {
		Name string `mapstructure:"name" json:"name"`
		Port int    `mapstructure:"port" json:"port"`
	}

	GrpcConfig struct {
		Ip   string `mapstructure:"ip" json:"ip"`
		Port int    `mapstructure:"port" json:"port"`
	}
)
