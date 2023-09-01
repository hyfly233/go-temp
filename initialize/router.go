package initialize

import (
	"go-temp/api"

	"github.com/gin-gonic/gin"
)

func InitGinEngine(_ *string) *gin.Engine {
	engine := gin.Default()

	engine.Use(gin.Recovery())

	// 添加路由
	group := engine.Group("/v1")
	api.ExampleRouter(group)

	return engine
}
