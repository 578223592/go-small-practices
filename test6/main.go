package main

import "fmt"

func main() {
	var nums = []int{}
	for i := 0; i < 1000; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)

	var stIds []string

	fmt.Println(stIds == nil)
	fmt.Println(len(stIds))
	for _, id := range stIds {
		fmt.Println(id)
	}
}
