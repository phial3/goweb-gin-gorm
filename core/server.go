package core

import (
	"fmt"
	"goweb-gin-gorm/cache"
	"goweb-gin-gorm/db"
	"goweb-gin-gorm/global"
	"time"
)

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

import (
	"goweb-gin-gorm/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	global.GlobalVp = Viper()          // 初始化Viper加载解析配置文件
	global.GlobalLog = InitZap()       // 初始化zap日志库
	global.GlobalDb = db.InitDb()      // 初始化 mysql DB
	global.GlobalRedis = cache.Redis() // 初始化 redis 配置

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GlobalConfig.System.Port)

	s := initServer(address, router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GlobalLog.Info("server run success on ", zap.String("address", address))
	global.GlobalLog.Error(s.ListenAndServe().Error())
}
