package cipher

import (
	"encoding/base64"
	"net/url"
)

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	resByte, err := base64.StdEncoding.DecodeString(s)
	return string(resByte), err
}

func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

func UrlDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}
