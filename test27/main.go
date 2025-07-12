package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	num := 0
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				num++
			})
		}()
	}
	time.Sleep(1) // 保证所有goroutine执行完毕
	fmt.Printf("num=%d", num)
}
