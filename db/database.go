package db

import (
	"go.uber.org/zap"
	"goweb-gin-gorm/global"
	"os"
)

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	m := global.GlobalConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		global.GlobalLog.Error("MySQL init error!", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// Database 在中间件中初始化mysql链接
func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GlobalConfig.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = global.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = global.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = global.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = global.Default.LogMode(logger.Info)
	default:
		config.Logger = global.Default.LogMode(logger.Info)
	}
	return config
}
