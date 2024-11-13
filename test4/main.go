package main

import (
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

}
