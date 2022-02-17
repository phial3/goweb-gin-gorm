package business

import (
	"github.com/gin-gonic/gin"
)

import (
	"goweb-gin-gorm/model"
	"goweb-gin-gorm/response"
)

// UserTokenRefreshService 用户刷新token的服务
type UserTokenRefreshService struct {
}

// Refresh 刷新token
func (service *UserTokenRefreshService) Refresh(c *gin.Context, user *model.User) response.Response {
	token, tokenExpire, err := user.MakeToken()
	if err != nil {
		response.ErrWithMessage("refresh token err", c)
	}
	data := model.BuildUser(*user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return response.Response{
		Data: data,
	}
}
