package main

import (
	"flag"
	"go-temp/initialize"
)

var Env *string

func main() {

	Env = flag.String("env", "dev", "程序的运行环境：dev(开发)、pro(生产)")
	flag.Parse()

	// 初始化
	initialize.InitLogger(Env)
	initialize.InitConf(Env)
	initialize.InitDB(Env)
}
