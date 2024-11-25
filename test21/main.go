package main

import (
	"log"
	"time"
)

// 读channel的多返回值中的ok表示是这个数据是否有效，而不是此时的channel是否关闭
func test1() {
	oneChan := make(chan int, 1)
	oneChan <- 1
	close(oneChan)
	val, ok := <-oneChan
	log.Printf("val:%d,ok:%v", val, ok) //2024/11/21 23:56:21 val:1,ok:true
}

// select过程中如果关闭chan
func test2() {
	oneChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		close(oneChan)
	}()

	select {
	case x, ok := <-oneChan:
		log.Printf("x:%d,ok:%v", x, ok)
	}
	log.Printf("ok")

}
func main() {
	// test1()
	test2()
}
