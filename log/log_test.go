package log

import (
	"fmt"
	"github.com/bang-go/kit/berror"
	"testing"
)

func TestLogger(t *testing.T) {
	log1, err := New(&Options{Default: ConfigDev})
	if err != nil {
		fmt.Println(err)
		berror.Exit(1)
	}
	log1.Info("hello")
	_, err = New(&Options{Default: ConfigProd})
	if err != nil {
		fmt.Println(err)
		berror.Exit(1)
	}
	//Error("test", String("say", "hello"))
	Info("hello")
}
