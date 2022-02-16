package user

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	baseRouterGroup := r.Group("base")
	// apiBaseGroup
	//var _ = api.ApiGroupApp.ApiBaseGroup
	{

	}
	return baseRouterGroup
}
