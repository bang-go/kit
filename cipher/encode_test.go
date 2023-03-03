package cipher

import (
	"fmt"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	s := "https://www.google.com/xxx"
	fmt.Println(UrlEncode(s))
	fmt.Println(UrlDecode(UrlEncode(s)))
}

func TestBase64Encode(t *testing.T) {
	s := "https://www.google.com/xxx"
	fmt.Println(Base64Encode(s))
	fmt.Println(Base64Decode(Base64Encode(s)))
}
