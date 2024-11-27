package main

import (
	"fmt"
	"time"
)

// 如何改变行为：修改go.mode文件中的go版本

func testFun() {

	// 程序退出信号
	quit := make(chan bool)
	timer := time.NewTimer(2 * time.Second)

	go func() {
		// 确保定时器已触发并发送信号
		time.Sleep(4 * time.Second)
		// 试图读取通道，看是否有值
		select {
		case t := <-timer.C:
			fmt.Println("接收到定时器信号：", t.Format(time.DateOnly))
		default:
			fmt.Println("无信号")
		}

		quit <- true
	}()

	// 确保定时器已触发并发送信号
	time.Sleep(3 * time.Second)
	wasStopped := timer.Stop()
	if wasStopped {
		// Go 1.23 或更高版本会走这条分支
		fmt.Println("定时器未过期，停止成功")
	} else {
		// Go 1.23 以前的版本会走这条分支
		fmt.Println("定时器已经过期并且信号已经发送")
	}

	// 等待退出信号
	<-quit
}
func main() {
	testFun()
}
