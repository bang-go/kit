package blog

import (
	"fmt"
	"github.com/bang-go/kit/berror"
	"testing"
)

func TestLog(t *testing.T) {
	log1, err := New(&Options{Default: ConfigProd})
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
	Info("hello1")
	log1.Info("hello2")
	_ = Configure(&Options{Default: ConfigDev})
	Info("hello1")
	log1.Info("hello2")

}
