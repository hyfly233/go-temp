package main

import (
	"context"
	"flag"
	"fmt"

	"os"
	"os/signal"
	"syscall"
	"time"

	"go-temp/constant"
	"go-temp/global"
	"go-temp/initialize"

	"go.uber.org/zap"
)

func main() {

	global.ENV = flag.String(constant.ENV, constant.DEV, fmt.Sprintf("程序的运行环境：%s(开发)、%s(生产)", constant.DEV, constant.PROD))
	flag.Parse()

	// 初始化
	initialize.InitLogger(global.ENV)
	initialize.InitConf(global.ENV)
	initialize.InitDB(global.ENV)

	srv := initialize.InitGinEngine(global.ENV)

	initialize.InitGrpc(global.ENV)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("关闭服务器中 ...")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Errorf("关闭服务器出错: %v", err)
	}

	select {
	case <-ctx.Done():
		zap.S().Info("timeout of 5 seconds.")
	}
	zap.S().Info("服务器终止")
}
