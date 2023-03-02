package bfloat

import (
	"github.com/bang-go/kit/constraint"
	"strconv"
)

// ToString FloatToString Float转化String
// prc precision精度
func ToString[T constraint.Float](val T, prc int) (str string) {
	str = strconv.FormatFloat(float64(val), 'f', prc, 64)
	return
}
