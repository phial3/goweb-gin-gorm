package util

import (
	"go.uber.org/zap"
	"math/rand"
	"os"
	"time"
)

// RandomString 返回随机字符串
func RandomString(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			Log().Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				Log().Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
