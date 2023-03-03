package util

import (
	"github.com/bang-go/kit/constraint"
)

// VarAddress 获取变量地址
func VarAddress[T constraint.Ordered](value T) *T {
	return &value
}

// VarValue VarAddress 获取变量值
func VarValue[T constraint.Ordered](value *T) T {
	return *value
}
