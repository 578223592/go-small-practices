package main

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		x += sum
		return x
	}
}
func main() {
	// 这个函数返回一个闭包函数
	var fun1 = adder()
	fun1(1)
}
