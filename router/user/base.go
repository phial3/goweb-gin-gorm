package user

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-gorm/api"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	baseRouterGroup := r.Group("base")
	// baseGropu
	var _ = api.ApiGroupApp.UserGroup
	{

	}
	return baseRouterGroup
}
