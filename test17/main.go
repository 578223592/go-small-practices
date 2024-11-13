// main.go
package main

import (
	"log"
	"sync"
	"sync/atomic"
)

//
//var mutex1 sync.Mutex //这么初始化没问题
//var mutex1Count atomic.Int64
//
//func funcMutex1() {
//	mutex1.Lock()
//	defer mutex1.Unlock()
//	mutex1Count.Add(1)
//}
//
//func cirFuncMutex1(num int) {
//	for i := 0; i < num; i++ {
//		funcMutex1()
//	}
//	log.Println(mutex1Count.Load())
//}

type User struct {
	Id int
}

// var mutex2 = sync.Mutex{} //这么初始化没问题
// var mutex2 = &sync.Mutex{} //这么初始化没问题
var mutex2 sync.Mutex //这么初始化没问题
//var mutex2 *sync.Mutex //这么初始化直接用会报空指针panic

var mutex2Count atomic.Int64

func funcMutex2() *User {
	mutex2.Lock()
	defer mutex2.Unlock()
	u := User{Id: 1}
	mutex2Count.Add(1)
	return &u
}

func cirFuncMutex2(num int) {
	for i := 0; i < num; i++ {
		funcMutex2()
	}
	log.Println(mutex2Count.Load())
}
func main() {
	//cirFuncMutex1(1000)
	cirFuncMutex2(1000)
}
