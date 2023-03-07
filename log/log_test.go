package log

import (
	"fmt"
	"github.com/bang-go/kit/berror"
	"testing"
)

func TestLogger(t *testing.T) {
	logger, err := New(&Options{Default: ConfigProd})
	if err != nil {
		fmt.Println(err)
		berror.Exit(1)
	}
	//Error("test", String("say", "hello"))
	logger.Info("hello")
}
