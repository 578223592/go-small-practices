package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	slice1 = slice1[0:5]
	fmt.Println(slice1)
	for i := 0; i < 5; i++ {
		slice1 = append(slice1, 0)
	}
	fmt.Println(slice1)
}
