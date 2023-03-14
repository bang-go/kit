package bhttp

import (
	"encoding/json"
	"net/url"
)

// FormatFormData 格式化form数据
func FormatFormData(data map[string]string) string {
	urlValues := url.Values{}
	for key, value := range data {
		urlValues.Set(key, value)
	}
	return urlValues.Encode()
}

// FormatJsonData 格式化json数据
func FormatJsonData(data map[string]string) string {
	byteBody, _ := json.Marshal(data)
	return string(byteBody)
}
