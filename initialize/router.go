package initialize

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-gorm/global"
	"goweb-gin-gorm/middleware"
	"goweb-gin-gorm/response"
	"goweb-gin-gorm/router"
	"os"
)

func Routers() *gin.Engine {
	var r = gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	routerGroup := router.RouterGroupApp

	// public
	publicGroup := r.Group("")
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) { response.Ok(c) })
	}
	{
		routerGroup.InitBaseRouter(publicGroup)
	}

	// private
	privateGroup := r.Group("")
	//路由要经过jwt和cas校验
	privateGroup.Use(middleware.AuthRequired()).Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		routerGroup.InitUserRouter(privateGroup)
	}

	global.GlobalLog.Info("router register success!")

	return r
}
