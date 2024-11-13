package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums = []int{}
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})
	fmt.Println(nums)
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	fmt.Println(nums)
}
