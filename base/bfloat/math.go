package bfloat

import (
	"github.com/bang-go/kit/constraint"
	"github.com/shopspring/decimal"
)

// MathAdd 加法
func MathAdd[T constraint.Float](args ...T) float64 {
	var result decimal.Decimal
	for _, v := range args {
		result = result.Add(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// MathSub 减法
func MathSub[T constraint.Float](minuend T, args ...T) float64 {
	result := decimal.NewFromFloat(float64(minuend))
	for _, v := range args {
		result = result.Sub(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// MathMul 乘法
func MathMul[T constraint.Float](arg1 T, args ...T) float64 {
	result := decimal.NewFromFloat(float64(arg1))
	for _, v := range args {
		result = result.Mul(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// MathDiv 除法
func MathDiv[T constraint.Float](dividend T, args ...T) float64 {
	result := decimal.NewFromFloat(float64(dividend))
	for _, v := range args {
		result = result.Div(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// MathCompare 比较
func MathCompare[T constraint.Float](v1 T, v2 T) int {
	cV1 := decimal.NewFromFloat(float64(v1))
	cV2 := decimal.NewFromFloat(float64(v2))
	return cV1.Cmp(cV2)
}

// MathAbs　绝对值
func MathAbs[T constraint.Float](v T) float64 {
	result := decimal.NewFromFloat(float64(v)).Abs()
	res, _ := result.Float64()
	return res
}
