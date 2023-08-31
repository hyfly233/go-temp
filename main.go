package main

import (
	"flag"
	"fmt"
	"go-temp/constant"
	"go-temp/initialize"
)

var Env *string

func main() {

	Env = flag.String(constant.ENV, constant.DEV, fmt.Sprintf("程序的运行环境：%s(开发)、%s(生产)", constant.DEV, constant.PROD))
	flag.Parse()

	// 初始化
	initialize.InitLogger(Env)
	initialize.InitConf(Env)
	initialize.InitDB(Env)
}
