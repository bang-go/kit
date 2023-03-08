package bgin

import (
	"testing"
)

func TestGin(t *testing.T) {
	client := New(&Options{Addr: ":8080"})
	gp := client.RouterGroup("/bot")
	gp.RouterHandle(MethodGet, "/health", func(c *Context) {
		c.JSON(200, "success")
	})
	client.Run()
}
