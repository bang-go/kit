package util

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	fmt.Println(GetDir())
	fmt.Println(GetFile())
	fmt.Println(GetLine())
}
