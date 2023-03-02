package env

import (
	"os"
)

// SetEnv 设置环境变量
func SetEnv(key string, value string) error {
	return os.Setenv(key, value)
}

// GetEnv 获取环境变量
func GetEnv(key string) string {
	return os.Getenv(key)
}
