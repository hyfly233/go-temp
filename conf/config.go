package conf

type (
	NacosConfig struct {
		Host      string `mapstructure:"host"`
		Port      uint64 `mapstructure:"port"`
		Namespace string `mapstructure:"namespace"`
		User      string `mapstructure:"user"`
		Password  string `mapstructure:"password"`
		DataId    string `mapstructure:"data_id"`
		Group     string `mapstructure:"group"`
	}

	ServerConfig struct {
		Name       string       `mapstructure:"name" json:"name"`
		Host       string       `mapstructure:"host" json:"host"`
		RestPort   int          `mapstructure:"rest_port" json:"rest_port"`
		RpcPort    int          `mapstructure:"rpc_port" json:"rpc_port"`
		Tags       []string     `mapstructure:"tags" json:"tags"`
		MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
		ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`

		// 其他 rpc 微服务的配置
		OtherSrvInfo OtherSrvConfig `mapstructure:"other_srv" json:"other_srv"`
	}

	MysqlConfig struct {
		Host     string `mapstructure:"host" json:"host"`
		Port     int    `mapstructure:"port" json:"port"`
		Name     string `mapstructure:"db" json:"db"`
		User     string `mapstructure:"user" json:"user"`
		Password string `mapstructure:"password" json:"password"`
	}

	ConsulConfig struct {
		Host string `mapstructure:"host" json:"host"`
		Port int    `mapstructure:"port" json:"port"`
	}

	OtherSrvConfig struct {
		Name string `mapstructure:"name" json:"name"`
	}
)
