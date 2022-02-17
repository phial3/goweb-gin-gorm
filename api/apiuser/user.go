package apiuser

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goweb-gin-gorm/model"
)

import (
	"goweb-gin-gorm/cache"
	"goweb-gin-gorm/response"
)

type UserApi struct {
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
func (u *UserApi) UserRegister(c *gin.Context) {
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		response.Result(res, c)
	} else {
		response.ErrorResponse(200, "", err)
	}
}

// UserLogin 用户登录接口
func (u *UserApi) UserLogin(c *gin.Context) {
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c)
		response.Result(res, c)
	} else {
		response.ErrorResponse(200, "", err)
	}
}

// UserTokenRefresh 用户刷新token接口
func (u *UserApi) UserTokenRefresh(c *gin.Context) {
	user := CurrentUser(c)
	res := userTokenRefreshService.Refresh(c, user)
	response.Result(res, c)
}

// UserDetail 用户详情
func (u *UserApi) UserDetail(c *gin.Context) {
	user := CurrentUser(c)
	res := model.BuildUserResponse(*user)
	response.Result(res, c)
}

// UserLogout 用户登出
func (u *UserApi) UserLogout(c *gin.Context) {
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
	response.Ok(c)
}
