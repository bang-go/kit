package redis

import (
	"fmt"
	"testing"
)

func TestGetKey(t *testing.T) {
	client := New(&Options{
		Config: RedisConfig{
			Addr:     "localhost:6379",
			Password: "test",
		},
	})
	res, err := client.Get("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
