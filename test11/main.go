package main

import (
	"fmt"
)

func modifySlice1(innerSlice []string) {
	innerSlice[0] = "b"
	innerSlice = append(innerSlice, "a") //在这里会触发扩容，所以这条及其之后的innerSlice和传入的slice不是同一个
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}

func modifySlice2(innerSlice []string) {
	innerSlice = append(innerSlice, "a") //由于传入的slice容量cap够，所以不会触发扩容，底层还是一个
	//！！但是由于len和cap本质上是slice的一个属性，所以修改了slice的len，不会影响到外面，外面的slice还是只打印两个元素
	/*
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}
func main() {
	outerSlice := []string{"a", "a123"}
	fmt.Println(outerSlice)

	modifySlice1(outerSlice)

	fmt.Println(outerSlice)

	fmt.Println("=======================")

	outerSlice = make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	fmt.Println(outerSlice)

	modifySlice2(outerSlice)

	fmt.Println(outerSlice) //虽然不会扩容，但是还是只会打印两个元素，具体原因见：modifySlice2
}
