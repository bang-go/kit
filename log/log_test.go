package log

import (
	"fmt"
	"github.com/bang-go/kit/error"
	"testing"
)

func TestLogger(t *testing.T) {
	err := InitLogger(&Options{Default: ConfigProd})
	if err != nil {
		fmt.Println(err)
		error.Exit(1)
	}
	Info("test", String("say", "hello"))
}
