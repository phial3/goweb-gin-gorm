package core

import (
	"flag"
	"fmt"
	"os"
	"time"
)

import (
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

import (
	"goweb-gin-gorm/constant"
	"goweb-gin-gorm/global"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			// 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
				config = constant.ConfigFile
				fmt.Printf("config path : %v\n", constant.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("config path : %v\n", config)
			}
		} else {
			fmt.Printf("config path : %v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("config path : %v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	// 加载读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控并重新读取配置文件的改动
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GlobalConfig); err != nil {
			fmt.Println(err)
		}
	})

	// 解析JSON配置
	if err := v.Unmarshal(&global.GlobalConfig); err != nil {
		fmt.Println(err)
	}

	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.GlobalConfig.JWT.ExpiresTime)),
	)
	return v
}
