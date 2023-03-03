package bint

import (
	"fmt"
	"testing"
)

func TestRandRange(t *testing.T) {
	var min int64 = 10
	var max int64 = 10000
	for i := 0; i <= 100; i++ {
		fmt.Println(RandRange(min, max))
	}
}
