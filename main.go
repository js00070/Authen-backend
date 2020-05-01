package main

import (
	"authen/conf"
	"authen/server"
)

func main() {
	// 导入配置文件
	conf.Init()

	r := server.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
