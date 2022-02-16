package api

import "goweb-gin-gorm/api/apiuser"

type ApiGroup struct {
	UserGroup apiuser.ApiUserGroup
}

var ApiGroupApp = new(ApiGroup)
