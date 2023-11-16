package initialize

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go-temp/api"
	"go-temp/global"
)

func InitGinEngine(_ *string) *http.Server {
	zap.S().Info("Gin 初始化 ------------------")

	engine := gin.Default()

	engine.Use(gin.Recovery())

	// 添加路由
	group := engine.Group("/v1")
	api.ExampleRouter(group)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", strconv.Itoa(global.ServerConfig.RestPort)),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorf("ginEngine 运行出错: %v", err)
		}
	}()

	zap.S().Info("Gin 初始化成功 ------------------")
	return srv
}
