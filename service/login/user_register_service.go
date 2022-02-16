package login

import (
	"goweb-gin-gorm/constant"
	"goweb-gin-gorm/global"
	"goweb-gin-gorm/model"
	"goweb-gin-gorm/response"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *response.Response {
	if service.PasswordConfirm != service.Password {
		return &response.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := int64(0)
	global.GlobalDb.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &response.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	global.GlobalDb.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &response.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() response.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   constant.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return response.Err(
			constant.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := global.GlobalDb.Create(&user).Error; err != nil {
		return response.ParamErr("注册失败", err)
	}
	return model.BuildUserResponse(user)
}
