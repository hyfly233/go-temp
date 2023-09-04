package conf

type (
	AppsConfig struct {
		Server struct {
			Port int    `mapstructure:"port"`
			Name string `mapstructure:"name"`
		} `mapstructure:"server"`
	}

	ServerConfig struct {
		Name string `mapstructure:"name" json:"name"`
		Port int    `mapstructure:"port" json:"port"`
	}
)
