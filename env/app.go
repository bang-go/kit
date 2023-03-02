package env

import (
	"fmt"
)

var (
	AppEnv string //应用环境变量
)

const (
	PROD             = "prod"    //生产环境
	GRAY             = "gray"    //灰度环境
	PRE              = "pre"     //预发布环境
	TEST             = "test"    //测试环境
	DEV              = "dev"     //开发环境
	DefaultAppEnvKey = "APP_ENV" //默认应用环境变量
)

var (
	AppEnvKey = DefaultAppEnvKey
)

// SetAppEnvKey SetAppEnv 设置环境变量
func SetAppEnvKey(key string) {
	AppEnvKey = key
}

// GetAppEnv 获取应用环境变量
func GetAppEnv() string {
	return AppEnv
}

// InitAppEnv 初始化应用环境变量
func InitAppEnv() {
	AppEnv = GetEnv(AppEnvKey)
	if len(AppEnv) == 0 { //默认开发环境
		AppEnv = DEV
	}
	switch AppEnv {
	case PROD, GRAY, PRE, TEST, DEV:
	default:
		panic(fmt.Sprintf("Unknown environment variable: %s", AppEnv))
	}
}

func IsProd() bool {
	if AppEnv == PROD {
		return true
	}
	return false
}
func IsDev() bool {
	if AppEnv == DEV {
		return true
	}
	return false
}
func IsTest() bool {
	if AppEnv == TEST {
		return true
	}
	return false
}

func IsPre() bool {
	if AppEnv == PRE {
		return true
	}
	return false
}

func IsGray() bool {
	if AppEnv == GRAY {
		return true
	}
	return false
}
