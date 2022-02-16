package main

import (
	"goweb-gin-gorm/conf"
	"goweb-gin-gorm/initialize"
)

func main() {
	// 从配置文件读取配置并加载服务（mysql）
	conf.Init()

	// 装载路由
	_ = initialize.Routers().Run(":3001")
}
