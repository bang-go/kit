package bslice

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{4, 5, 6, 7, 8, 8, 1, 9}
	Sort(s)
	fmt.Println(s)
}
