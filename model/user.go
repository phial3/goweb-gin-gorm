package model

import (
	"goweb-gin-gorm/constant"
	"goweb-gin-gorm/db"
	"goweb-gin-gorm/response"
	"strconv"
	"time"
)

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

import (
	"goweb-gin-gorm/cache"
	"goweb-gin-gorm/util"
)

// User 用户模型
type User struct {
	gorm.Model
	ID          uint   `json:"ID" gorm:"primarykey"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	Nickname    string `json:"nickname"`
	Status      string `json:"status"`
	Avatar      string `gorm:"size:1000"`
	Token       string `json:"token,omitempty"`
	TokenExpire int64  `json:"token_expire,omitempty"`
	CreatedAt   int64
}

// BuildUser 序列化用户
func BuildUser(user User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user User) response.Response {
	return response.Response{
		Data: BuildUser(user),
	}
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := db.DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), constant.PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// UserID 返回string版的uid
func (user *User) UserID() string {
	return strconv.Itoa(int(user.ID))
}

// MakeToken 生成token
func (user *User) MakeToken() (string, int64, error) {
	// 移动端生成token, 2周自动过期
	token := util.RandStringRunes(15)
	exp := 14 * 24 * time.Hour
	tokenExpire := time.Now().Add(exp).Unix()
	if err := cache.SaveUserToken(token, user.UserID(), exp); err != nil {
		return "", 0, err
	}
	return token, tokenExpire, nil
}
