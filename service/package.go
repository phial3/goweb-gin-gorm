package service

import (
	"goweb-gin-gorm/service/business"
	"goweb-gin-gorm/service/login"
)

type ServiceGroup struct {
	LoginGroup    login.LoginGroup
	BusinessGroup business.BusinessGroup
}

var ServiceGroupApp = new(ServiceGroup)
