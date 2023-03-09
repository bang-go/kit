package goroutine

import "github.com/bang-go/kit/berror"

type GoFunc func()

func New(fc GoFunc) {
	go func() {
		defer berror.PanicRecover()
		fc()
	}()
}
