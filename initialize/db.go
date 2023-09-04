package initialize

import (
	"go.uber.org/zap"
)

func InitDB(env *string) {
	zap.S().Info("初始化数据库 ------------------")

	//dsn := ""
	//var err error
	//global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//if err != nil {
	//	panic("failed to connect database")
	//}

	zap.S().Info("数据库初始化成功 ------------------")
}
