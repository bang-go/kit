package utils

import "github.com/google/uuid"

// NewUUID 生成uuid字符串
func NewUUID() string {
	return uuid.New().String()
}
