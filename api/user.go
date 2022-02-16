package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goweb-gin-gorm/model"
)

import (
	"goweb-gin-gorm/cache"
	"goweb-gin-gorm/response"
	"goweb-gin-gorm/service"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "Ping",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var registerService service.UserRegisterService
	if err := c.ShouldBind(&registerService); err == nil {
		res := registerService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var loginService service.UserLoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// UserTokenRefresh 用户刷新token接口
func UserTokenRefresh(c *gin.Context) {
	user := CurrentUser(c)
	var refreshService service.UserTokenRefreshService
	res := refreshService.Refresh(c, user)
	c.JSON(200, res)
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := model.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	// 移动端登出
	token := c.GetHeader("X-Token")
	if token != "" {
		_ = cache.RemoveUserToken(token)
	} else {
		// web端登出
		s := sessions.Default(c)
		s.Clear()
		s.Save()
	}
	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
