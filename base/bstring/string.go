package bstring

import (
	"strconv"
	"strings"
)

// Int StringToInt  String转化为Int
func Int(val string) (res int, err error) {
	return strconv.Atoi(val)
}

// Float StringToFloat  String转化为Float
func Float(val string) (res float64, err error) {
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
