package cache

import (
	"github.com/go-redis/redis/v8"
)

import (
	"goweb-gin-gorm/global"
)

// Redis 在中间件中初始化redis链接
func Redis() *redis.Client {
	redisConf := global.GlobalConfig.Redis

	client := redis.NewClient(&redis.Options{
		Addr:       redisConf.Addr,
		Password:   redisConf.Password,
		DB:         redisConf.DB,
		MaxRetries: 1,
	})

	return client
}
