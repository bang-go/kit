package bint

import (
	"github.com/bang-go/kit/constraint"
	"strconv"
)

// ToString Int转化String(10进制)
func ToString[T constraint.Integer](val T) (str string) {
	str = strconv.FormatInt(int64(val), 10)
	return
}

// IsEmpty 是否为0
func IsEmpty[T constraint.Integer](val T) bool {
	if val == 0 {
		return true
	}
	return false
}
