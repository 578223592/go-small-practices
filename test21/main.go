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
		// time.Sleep(time.Second)
		close(oneChan)
	}()

	select {
	case x, ok := <-oneChan:
		log.Printf("x:%d,ok:%v", x, ok)
	}
	// log.Printf("ok")

}

func test3() {
	oneChan := make(chan int, 10)
	go func() {
		time.Sleep(time.Second)
		close(oneChan)
	}()

	oneChan <- 1
}

func test4() {
	go func() {
		for {
			var i = 1
			for i := 0; i < 100; i++ {
				i++
			}
			log.Print(i)
		}
	}()
	oneChan := make(chan int, 10)
	oneChan = nil
	oneChan <- 1

}
func main() {
	//test1()
	// test2()
	//test3()
	test4()
}
