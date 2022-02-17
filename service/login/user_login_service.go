package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goweb-gin-gorm/global"
)

import (
	"goweb-gin-gorm/model"
	"goweb-gin-gorm/response"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Token    bool   `form:"token" json:"token"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) response.Response {
	var user model.User

	if err := global.GlobalDb.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return response.ParamErr(-1, "账号或密码错误", nil)
	}

	if !user.CheckPassword(service.Password) {
		return response.ParamErr(-1, "账号或密码错误", nil)
	}

	var token string
	var tokenExpire int64
	var err error
	if service.Token {
		token, tokenExpire, err = user.MakeToken()
		if err != nil {
			response.ErrWithMessage("get token err", c)
		}
	} else {
		// web端设置session
		service.setSession(c, user)
	}

	data := model.BuildUser(user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return response.Response{
		Data: data,
	}
}

// Logout 注销操作，清除session
func (service *UserLoginService) Logout(c *gin.Context) response.Response {
	return response.Response{}
}
