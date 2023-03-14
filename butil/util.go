package butil

import "github.com/bang-go/kit/constraint"

// If 三元表达式
func If[T constraint.Ordered](expr bool, trueVale, falseValue T) T {
	if expr {
		return trueVale
	}
	return falseValue
}
