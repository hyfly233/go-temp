package api

import (
	"net/http"

	"go-temp/global"
	"go-temp/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ExampleRouter(group *gin.RouterGroup) {
	zap.S().Info("添加 Example Router ------------------")
	{
		group.GET("/test01", getTest01)
		group.GET("/test02/:id", getTest02)

		group.POST("/test03", postTest03)
	}
}

func getTest01(ctx *gin.Context) {
	test := new(model.Test)
	tx := global.DB.Table("test").Where("id = ?", 1).First(&test)

	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "查询出错，错误详情：" + tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "test01" + test.Id,
	})
}

func getTest02(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "test02",
		"id":      id,
	})
}

func postTest03(ctx *gin.Context) {
	m := new(model.Test)
	err := ctx.BindJSON(&m)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误，错误详情：" + err.Error(),
		})
		return
	}

	id := m.Id
	zap.S().Info("postTest03 id: ", id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "test03",
		"id":      id,
	})
}
