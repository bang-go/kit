package bint

import (
	"github.com/bang-go/kit/constraint"
	"math/rand"
	"time"
)

// RandRange 区间随机数
func RandRange[T constraint.Integer](n1 T, n2 T) int64 {
	if n2 < n1 {
		n1, n2 = n2, n1
	}
	if n1 == n2 {
		return int64(n1)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(int64(n2-n1)) + int64(n1)
}
