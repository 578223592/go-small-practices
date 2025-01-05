package main

import (
	"sync/atomic"
	"time"
)

var counter int32 = 1

func atomicWriter() {
	for {
		atomic.AddInt32(&counter, 1) // 原子写操作

	}
}
func writer() {
	for {
		counter += 1
		//time.Sleep(1 * time.Millisecond)
	}
}

func reader() {
	for {
		_ = counter
		//_ = counter
		//_ = counter
		//_ = counter
		//_ = counter
		//_ = counter
		//fmt.Println("Current counter:", counter) // 直接读读操作
		//time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	//go atomicWriter()
	//go writer()
	go reader()

	time.Sleep(3 * time.Second)
}
