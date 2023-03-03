package util

import (
	"fmt"
	"testing"
)

func TestVarAddress(t *testing.T) {
	var a int
	val := VarAddress(a)
	fmt.Println(val)
}

func TestVarValue(t *testing.T) {
	type num int
	var a num
	val := VarValue(&a)
	fmt.Println(val)
}
