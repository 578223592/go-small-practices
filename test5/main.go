package main

import (
	"fmt"
	"math"
	"sync"
)
import jsoniter "github.com/json-iterator/go"

// JSON 比原生json性能更好的解析器，用法与原生json一样
// ddjson.JSON.Marshal()
// ddjson.JSON.Unmarshal()
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

func NewStudent() *Student {
	return &Student{}
}

type Student struct {
	Name  string `json:"name,omitempty"`
	Age   int
	Right bool
}

func NewWithoutPool(num int) {
	for i := 0; i < num; i++ {
		NewStudent()
	}
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return NewStudent()
	},
}

func NewWithPool(num int) {

}

func main() {

	fmt.Printf("int max:%d\n", math.MaxInt)
	fmt.Printf("int max + 1 to uint:%d\n", uint(math.MaxInt+1))
	fmt.Printf("int max to uint:%d\n", uint(math.MaxInt))
	//fmt.Printf("int min to uint:%d\n", uint(math.MinInt)) // constant -9223372036854775808 overflows uint

	var numInt int = -1
	var minInt int = math.MinInt
	fmt.Printf("%d \n", uint(minInt))

	fmt.Printf("-1 to uint %d \n", uint(numInt))

	fmt.Printf("%d \n", minInt-1) // 应该等于math.MaxInt，具体参考：计算机正码、反码、补码 https://www.cnblogs.com/zhangziqiu/archive/2011/03/30/ComputerCode.html

}
