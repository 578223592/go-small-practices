package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"math/rand"
)

func main() {
	slices := make([]int, 100)
	fmt.Println(slices)
	ints := slices[0:100]
	fmt.Println(ints)
	//fmt.Println(slices[1:102]) //panic: runtime error: slice bounds out of range [:102] with capacity 100
	//fmt.Println(slices[105:105]) // panic: runtime error: slice bounds out of range [:105] with capacity 100
	//for i := 0; i < 1000; i++ {
	//	i2 := rand.Int()
	//	fmt.Println(slices[i2:10]) // panic: runtime error: slice bounds out of range [237517091944302022:10]
	//}

	for i := 0; i < 10000; i++ { //测试一下limit+offset 的健壮性
		currentIndex := rand.Int()
		pageSize := rand.Int()
		totalCount := len(slices)
		if currentIndex <= 1 {
			continue
		}
		if pageSize <= 0 {
			continue
		}
		_ = slices[cast.ToInt64(math.Min(float64(currentIndex*pageSize), float64(totalCount))):cast.ToInt64(math.Min(float64(currentIndex*pageSize+pageSize), float64(totalCount)))]

	}
}
