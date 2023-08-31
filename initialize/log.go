package initialize

import (
	"fmt"

	"go.uber.org/zap"
)

func InitLogger(env *string) {

	//Env = flag.String("env", "dev", "程序的运行环境：dev(开发)、pro(生产)")
	//flag.Parse()

	var (
		logger *zap.Logger
		err    error
	)

	switch *env {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "pro":
		logger, err = zap.NewProduction()
	default:
		panic(fmt.Sprintf("初始化日志出错,不支持的 env 参数\n"))
	}

	if err != nil {
		panic(fmt.Sprintf("初始化日志出错,错误详情：%s\n", err.Error()))
	}

	zap.ReplaceGlobals(logger)
	zap.S().Info("日志初始化成功------------------")
}
