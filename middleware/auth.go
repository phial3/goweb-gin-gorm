package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goweb-gin-gorm/response"
)

import (
	"goweb-gin-gorm/cache"
	"goweb-gin-gorm/model"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid string
		token := c.GetHeader("X-Token")
		if token != "" {
			uid, _ = cache.GetUserByToken(token)
		} else {
			session := sessions.Default(c)
			uid, _ = session.Get("user_id").(string)
		}
		if uid != "" {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		response.OkWithMessage("未登录", c)
		c.Abort()
	}
}
