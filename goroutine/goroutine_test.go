package goroutine

import (
	"errors"
	"fmt"
	"github.com/bang-go/kit/berror"
	"testing"
)

func TestGo(t *testing.T) {
	defer berror.PanicRecover()
	New(func() {
		a := func() error {
			panic("panic xxxx")
			return errors.New("eee")
		}
		if err := a(); err != nil {
			fmt.Println(err)
		}
	})
	panic("yyy")
	fmt.Println("end")
}
