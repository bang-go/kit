package http

import (
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	res := &Request{
		Method: MethodPost,
		Url:    "https://api.xxx.com",
		Form:   map[string]string{"world": "1"},
		Cookies: map[string]string{
			"token": "xxx",
		},
	}

	resp, err := res.Send()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp)
	fmt.Println(string(resp.Content))
	//fmt.Println(resp.Cookies)
}
