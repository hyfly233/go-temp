package initialize

import (
	"fmt"

	"go-temp/constant"

	"go.uber.org/zap"
)

func InitLogger(env *string) {

	var (
		logger *zap.Logger
		err    error
	)

	switch *env {
	case constant.DEV:
		logger, err = zap.NewDevelopment()
	case constant.PROD:
		logger, err = zap.NewProduction()
	default:
		panic(fmt.Sprintf("初始化日志出错,不支持的 env 参数\n"))
	}

	if err != nil {
		panic(fmt.Sprintf("初始化日志出错,错误详情：%s\n", err.Error()))
	}

	zap.ReplaceGlobals(logger)
	zap.S().Info("日志初始化成功 ------------------")
}
