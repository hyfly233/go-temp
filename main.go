package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"strconv"
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

	ginEngine := initialize.InitGinEngine(global.ENV)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", strconv.Itoa(initialize.AppConfig.Server.Port)),
		Handler: ginEngine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorf("ginEngine 运行出错: %v", err)
		}
	}()

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
