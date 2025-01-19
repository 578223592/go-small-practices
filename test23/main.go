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

// FnTimer1 Timer的使用用法
func FnTimer1() {
	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("timer 的启动时间为:%v\n", time.Now())

	expire := <-timer.C
	fmt.Printf("timer 的触发时间为:%v\n", expire)
}

func FnTimer2() {
	timeChannel := time.After(time.Second * 5)
	fmt.Printf("timer 的启动时间为:%v\n", time.Now())

	expire := <-timeChannel
	fmt.Printf("timer 的触发时间为:%v\n", expire)
}

func FnTimer3() {
	_ = time.AfterFunc(time.Second*5, func() {
		fmt.Printf("定时器触发了,触发时间为%v\n", time.Now())
	})

	fmt.Printf("timer 的启动时间为:%v\n", time.Now())
	time.Sleep(10 * time.Second) // 确保定时器触发
}

// FnTicket1 Ticker的错误的使用方法！！
func FnTicket1() {
	ticker := time.NewTicker(1 * time.Second)
	//defer ticker.Stop()
	go func() {
		for now := range ticker.C {
			// do something
			fmt.Printf("ticker 的触发时间为:%v\n", now)
		}
		fmt.Println("ticker 结束")
	}()

	time.Sleep(4 * time.Second)
	ticker.Stop()
}

// FnTicket2 Ticker的正确的使用方法！！
func FnTicket2() {
	ticker := time.NewTicker(1 * time.Second)
	stopTicker := make(chan struct{})
	defer ticker.Stop()
	go func() {
		for {
			select {
			case now := <-ticker.C:
				// do something
				fmt.Printf("ticker 的触发时间为:%v\n", now)
			case <-stopTicker:
				fmt.Println("ticker 结束")
				return
			}
		}
	}()

	time.Sleep(4 * time.Second)
	stopTicker <- struct{}{}
}

func SomeTry() {
	go func() {
		for {
			func() {
				timer := time.NewTimer(time.Second * 2)
				defer timer.Stop()

				select {
				case b := <-c:
					if !b {
						fmt.Println(time.Now(), "work...")
					}
				case <-timer.C: // BBB: normal receive from channel timeout event
					fmt.Println(time.Now(), "timeout")
				}
			}()
		}
	}()
}

func main() {
	//FnTimer1()
	//FnTimer2()
	//FnTimer3()
	//FnTicket1()
	FnTicket2()
	//testFun()
	time.Sleep(999 * time.Second) //保证main不退出
}
