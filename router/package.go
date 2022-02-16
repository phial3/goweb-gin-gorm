package router

import (
	"goweb-gin-gorm/router/user"
)

type RouterGroup struct {
	user.BaseRouter
	user.UserRouter
}

var RouterGroupApp = new(RouterGroup)
