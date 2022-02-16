package util

import (
	"fmt"
	"goweb-gin-gorm/constant"
	"os"
	"time"
)

var logger *Logger

// Logger 日志
type Logger struct {
	level int
}

// Println 打印
func (ll *Logger) Println(msg string) {
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
}

// Panic 极端错误
func (ll *Logger) Panic(format string, v ...interface{}) {
	if constant.LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Panic] "+format, v...)
	ll.Println(msg)
	os.Exit(0)
}

// Error 错误
func (ll *Logger) Error(format string, v ...interface{}) {
	if constant.LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	ll.Println(msg)
}

// Warning 警告
func (ll *Logger) Warning(format string, v ...interface{}) {
	if constant.LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	ll.Println(msg)
}

// Info 信息
func (ll *Logger) Info(format string, v ...interface{}) {
	if constant.LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	ll.Println(msg)
}

// Debug 校验
func (ll *Logger) Debug(format string, v ...interface{}) {
	if constant.LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf("[D] "+format, v...)
	ll.Println(msg)
}

// BuildLogger 构建logger
func BuildLogger(level string) {
	intLevel := constant.LevelError
	switch level {
	case "error":
		intLevel = constant.LevelError
	case "warning":
		intLevel = constant.LevelWarning
	case "info":
		intLevel = constant.LevelInformational
	case "debug":
		intLevel = constant.LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

// Log 返回日志对象
func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: constant.LevelDebug,
		}
		logger = &l
	}
	return logger
}
