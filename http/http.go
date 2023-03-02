package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	//请求方式
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
	//状态码
	StatusOk = http.StatusOK
)

// Request 请求结构体
type Request struct {
	Url      string            `json:"url"`      // 请求url
	Method   string            `json:"method"`   //请求方法，GET/POST/PUT/DELETE/PATCH...
	Headers  map[string]string `json:"headers"`  // 请求头
	Params   map[string]string `json:"params"`   //Query参数
	Raw      string            `json:"raw"`      //原始请求数据
	Form     map[string]string `json:"form"`     //表单格式请求数据
	Json     map[string]string `json:"json"`     //JSON格式请求数据
	Files    map[string]string `json:"files"`    //TODO:文件
	Cookies  map[string]string `json:"cookies"`  //Cookies
	Timeout  time.Duration     `json:"timeout"`  //超时时间
	Username string            `json:"username"` //base认证username
	Password string            `json:"password"` //base认证password
	//TODO:代理
}

// Response 响应结构体
type Response struct {
	StatusCode int               `json:"status_code"` //状态码
	Success    bool              `json:"success"`     //响应状态
	Content    []byte            `json:"content"`     //响应内容-字节
	Reason     string            `json:"reason"`      // 状态码说明
	Elapsed    float64           `json:"elapsed"`     // 请求耗时(秒)
	Headers    map[string]string `json:"headers"`     // 响应头
	Cookies    map[string]string `json:"cookies"`     // 响应Cookies
	Request    *Request          `json:"request"`     // 原始请求
}

var (
	DefaultTimeout time.Duration = 30 * time.Second //默认请求时间
)

// 设置请求方式
func (r *Request) getMethod() (method string, err error) {
	if r.Method == "" {
		err = errors.New("method is empty")
		return
	}
	method = strings.ToUpper(r.Method)
	switch method {
	case MethodGet, MethodHead, MethodPost, MethodPut, MethodPatch, MethodDelete, MethodConnect, MethodOptions, MethodTrace:
	default:
		err = errors.New("unknown method")
		return
	}
	return
}

// 设置请求地址(拼接请求参数)
func (r *Request) getUrl() (us string, err error) {
	us = r.Url
	if r.Url == "" {
		err = errors.New("url is empty")
		return
	}
	if r.Params == nil {
		return
	}
	urlValues := url.Values{}
	httpUrl, err := url.Parse(r.Url)
	for key, value := range r.Params {
		urlValues.Set(key, value)
	}
	httpUrl.RawQuery = urlValues.Encode()
	us = httpUrl.String()
	return
}

// 设置请求头
func (r *Request) initHeaders(req *http.Request) {
	if r.Headers == nil {
		return
	}
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}
}

// 设置base认证
func (r *Request) initBaseAuth(req *http.Request) {
	if r.Username != "" && r.Password != "" {
		req.SetBasicAuth(r.Username, r.Password)
	}
}

func (r *Request) initCookie(req *http.Request) {
	if r.Cookies == nil {
		return
	}
	for key, value := range r.Cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}
}

// 设置超时时间
func (r *Request) initTimeout(client *http.Client) {
	if r.Timeout != 0 {
		client.Timeout = r.Timeout
		return
	}
	client.Timeout = DefaultTimeout
}

// 设置请求体
func (r *Request) getBody() *strings.Reader {
	var body string
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
	if len(r.Raw) > 0 { //原始数据
		body = r.Raw
	} else if r.Form != nil {
		urlValues := url.Values{}
		for key, value := range r.Form {
			urlValues.Set(key, value)
		}
		body = urlValues.Encode()
		r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	} else if r.Json != nil {
		byteBody, _ := json.Marshal(r.Json)
		body = string(byteBody)
		r.Headers["Content-Type"] = "application/json"
	}
	return strings.NewReader(body)
}

// Send 发送请求
func (r *Request) Send() (resp *Response, err error) {
	httpUrl, err := r.getUrl()
	if err != nil {
		return
	}
	method, err := r.getMethod()
	if err != nil {
		return
	}
	reqBody := r.getBody()
	var httpReq *http.Request
	var httpRes *http.Response
	if httpReq, err = http.NewRequest(method, httpUrl, reqBody); err != nil { //新建http请求
		return
	}

	r.initHeaders(httpReq)
	r.initBaseAuth(httpReq)
	r.initCookie(httpReq)
	client := &http.Client{}
	r.initTimeout(client)

	startTime := time.Now()
	if httpRes, err = client.Do(httpReq); err != nil {
		return
	}
	defer httpRes.Body.Close()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Seconds()
	resp = r.makeResponse(httpRes, elapsed)
	return
}

// 组装响应体
func (r *Request) makeResponse(res *http.Response, elapsed float64) (response *Response) {
	response = &Response{
		Request:    r,
		Elapsed:    elapsed,
		StatusCode: res.StatusCode,
	}
	response.Content, _ = io.ReadAll(res.Body)
	response.Reason = strings.Split(res.Status, " ")[1]
	response.Headers = map[string]string{}
	response.Cookies = map[string]string{}
	if res.StatusCode == StatusOk { //200 定义成功
		response.Success = true
	}
	for key, value := range res.Header {
		response.Headers[key] = strings.Join(value, ";")
	}
	for _, v := range res.Cookies() {
		response.Cookies[v.Name] = v.Value
	}
	return
}
