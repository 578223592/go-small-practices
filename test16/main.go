// main.go
package main

import "log"

type MyHandlers struct {
	index    int
	funcList []func(*MyHandlers) //关键点：函数的列表类型中入参为 *MyHandlers。1.自身；2.指针
}

func NewMyHandlers() *MyHandlers {
	m := new(MyHandlers)
	m.index = -1 //小细节：设置初始值=-1
	return m
}
func (h *MyHandlers) Add(f func(*MyHandlers)) {
	h.funcList = append(h.funcList, f)
}

// HandleNext 开始执行，可以看出是类似于链表
func (h *MyHandlers) HandleNext() {
	h.index++
	for ; h.index < len(h.funcList); h.index++ {
		h.funcList[h.index](h)
	}
}

func main() {

	myMiddleWare := NewMyHandlers()
	myMiddleWare.Add(func1)
	myMiddleWare.Add(fun2)
	myMiddleWare.HandleNext()
	/** func1 -> func2 ->func1
	2024/10/23 19:59:37 this is fun1 start
	2024/10/23 19:59:37 this is fun2 start
	2024/10/23 19:59:37 this is fun2 end
	2024/10/23 19:59:37 this is fun1 end
	*/
}

func func1(h *MyHandlers) {
	log.Print("this is fun1 start")
	h.HandleNext()
	log.Print("this is fun1 end")
}

func fun2(h *MyHandlers) {
	log.Print("this is fun2 start")
	log.Print("this is fun2 end")
}
