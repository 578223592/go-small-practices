// main.go
package main

import (
	"log"
)

func main() {
	{
		func() {
			defer log.Printf("执行defer")
			panic("123")

		}()
	}
	log.Printf("执行...")
}

/**
{
	defer log.Printf("执行defer")
}
log.Printf("执行...")


打印结果：
2024/10/18 10:24:05 执行...
2024/10/18 10:24:05 执行defer

*/
