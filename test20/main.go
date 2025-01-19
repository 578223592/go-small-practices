package main

import (
	"net/http"
)

// chainBuildExample
// @Description: 链式调用实现的核心思路：返回本身的指针以支持下一个调用；注意是 指针
func chainBuildExample() {
	// 使用链式调用创建一个GET请求
	request, _ := http.NewRequest("GET", "https://www.example.com", nil)
	req := &Request{request}
	req.WithHeader("User-Agent", "MyCustomUserAgent").
		WithCookie(&http.Cookie{Name: "session", Value: "abc123"})
}

type Request struct {
	*http.Request
}

// WithHeader 添加请求头
func (r *Request) WithHeader(key, value string) *Request {
	r.Header.Add(key, value)
	return r
}

// WithCookie 添加Cookie
func (r *Request) WithCookie(cookie *http.Cookie) *Request {
	r.AddCookie(cookie)
	return r
}

func main() {
	targetStruct := NewTargetStruct(withA(15)) //传入withA(15)就设置A，不传就不设置
	println(targetStruct.a)
}

type TargetStruct struct {
	a int
}

func NewTargetStruct(diffrentSetting ...func(t *TargetStruct) *TargetStruct) *TargetStruct {
	t := &TargetStruct{}
	for _, fun := range diffrentSetting {
		t = fun(t)
	}
	return t
}

func withA(a int) func(t *TargetStruct) *TargetStruct {
	return func(t *TargetStruct) *TargetStruct {
		t.a = a
		return t
	}
}
