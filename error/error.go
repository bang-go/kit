package error

import (
	"errors"
	"os"
)

func Exit(code int) {
	os.Exit(code)
}

func New(text string) error {
	return errors.New(text)
}
