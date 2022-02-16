package cache

import (
	"goweb-gin-gorm/global"
	"time"
)

// 把用户token保存在redis中
func SaveUserToken(token, userID string, expTime time.Duration) error {
	var err = global.GlobalRedis.Set(global.GlobalRedis.Context(), UserTokenKey(token), userID, expTime).Err()
	return err
}

// 通过Token获取对应的userid
func GetUserByToken(token string) (string, error) {
	return global.GlobalRedis.Get(global.GlobalRedis.Context(), UserTokenKey(token)).Result()
}

// 删除用户登录token实现登出
func RemoveUserToken(token string) error {
	return global.GlobalRedis.Del(global.GlobalRedis.Context(), UserTokenKey(token)).Err()
}
