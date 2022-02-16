package apiuser

import "goweb-gin-gorm/service"

type ApiUserGroup struct {
	UserApi
}

var userLoginService = service.ServiceGroupApp.LoginGroup.UserLoginService
var userRegisterService = service.ServiceGroupApp.LoginGroup.UserRegisterService
var userTokenRefreshService = service.ServiceGroupApp.BusinessGroup.UserTokenRefreshService
