package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"goweb-gin-gorm/config"
	timer "goweb-gin-gorm/util"

	"golang.org/x/sync/singleflight"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GlobalVp     *viper.Viper
	GlobalConfig config.Server

	GlobalDb    *gorm.DB
	GlobalRedis *redis.Client

	GlobalLog                *zap.Logger
	GlobalTimer              timer.Timer = timer.NewTimerTask()
	GlobalConcurrencyControl             = &singleflight.Group{}

	BlackCache local_cache.Cache
)
