package main

import (
	"fmt"
	"unsafe"
)

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	bytes := []byte{'a', 'b'}
	stringFromBytes1 := string(bytes)
	stringFromBytes2 := Bytes2String(bytes)
	stringFromBytes3 := stringFromBytes2

	bytes[0] = 'b'

	fmt.Println(stringFromBytes1)
	fmt.Println(stringFromBytes2) //由于采用了指针强转的方式，因此本质上这个字符串底层和[]byte是共享的，[]byte的修改也会影响这里。
	fmt.Println(stringFromBytes3)
	/* 打印结果：
	ab
	bb
	bb
	*/
}
