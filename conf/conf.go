package conf

import (
	"goweb-gin-gorm/db"
	"goweb-gin-gorm/model"
	"goweb-gin-gorm/response"
	"os"
)

import (
	"github.com/joho/godotenv"
)

import (
	"goweb-gin-gorm/cache"
	"goweb-gin-gorm/util"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := response.LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库, 数据自动迁移模式
	err := db.Database(os.Getenv("MYSQL_DSN")).AutoMigrate(&model.User{})
	if err != nil {
		panic("init database err!" + err.Error())
	}

	// 缓存
	cache.Redis()
}
