package bstring

import (
	"strconv"
	"strings"
)

// ToInt StringToInt  String转化为Int
func ToInt(val string) (res int, err error) {
	return strconv.Atoi(val)
}

// ToFloat StringToFloat  String转化为Float
func ToFloat(val string) (res float64, err error) {
	return strconv.ParseFloat(val, 64)
}

// ContainValue string 是否包含目标值
func ContainValue(str string, substr string) bool {
	return strings.Contains(str, substr)
}

// IsEmpty 是否为空
func IsEmpty(val string) bool {
	if val == "" {
		return true
	}
	return false
}
