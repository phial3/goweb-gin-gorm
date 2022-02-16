package user

import (
	"goweb-gin-gorm/api"
)

import (
	"github.com/gin-gonic/gin"
)

import (
	"goweb-gin-gorm/middleware"
)

type UserRouter struct {
}

// InitUserRouter UserRouter 路由配置
func (s *UserRouter) InitUserRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	userGroup := r.Group("user").Use(middleware.AuthRequired())
	var userApi = api.ApiGroupApp.UserGroup.UserApi
	{
		// 用户注册
		userGroup.POST("register", userApi.UserRegister)
		// 用户登录
		userGroup.POST("login", userApi.UserLogin)
		// User Routing
		userGroup.GET("detail", userApi.UserDetail)
		userGroup.DELETE("logout", userApi.UserLogout)
		// 用户刷新token
		userGroup.POST("refresh", userApi.UserTokenRefresh)
	}
	return userGroup
}
