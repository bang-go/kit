package cipher

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	str := "hello world"
	fmt.Println(Sha256(str))
}

func TestHmac(t *testing.T) {
	str := "hello world"
	key := "123"
	fmt.Println(Hmac(key, str))
}

func BenchmarkSha256(b *testing.B) {
	str := "hello world"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sha256(str)
	}
}
