package utils

import (
	"path"
	"runtime"
)

// GetDir 获取调用者目录
func GetDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// GetFile 获取调用者文件地址
func GetFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// GetLine 获取调用者行数
func GetLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line

}
