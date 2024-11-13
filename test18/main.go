// main.go
package main

import "log"

/*
对于嵌套实现，需要充分利用到go语言中将函数作为一等公名的必包特性
*/
type middleware func(next func()) func()

type MyHandlers struct {
	middlewares []middleware
	doFunc      func()
}

func (h *MyHandlers) Add(f middleware) {
	h.middlewares = append(h.middlewares, f)
}

func (h *MyHandlers) PrePareDo(doFunc func()) {
	//prePare
	for i := len(h.middlewares) - 1; i >= 0; i-- { //越早注册的包裹在越外层，这样可以越早执行
		doFunc = h.middlewares[i](doFunc)
	}
	//do
	doFunc()
	/* 输出
	2024/10/24 12:57:44 func1 start
	2024/10/24 12:57:44 func2 start
	2024/10/24 12:57:44 这里是业务代码
	2024/10/24 12:57:44 func2 end
	2024/10/24 12:57:44 func1 end
	*/
}

func main() {
	myHandlers := MyHandlers{}
	myHandlers.Add(warpFunc1)
	myHandlers.Add(warpFunc2)
	myHandlers.PrePareDo(someFunc)
}

func warpFunc1(next func()) func() {
	return func() {
		log.Println("func1 start")
		next()
		log.Println("func1 end")
	}
}

func warpFunc2(next func()) func() {
	return func() {
		log.Println("func2 start")
		next()
		log.Println("func2 end")
	}
}

// 业务代码
func someFunc() {
	log.Println("这里是业务代码")
}
