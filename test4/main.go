package main

import (
	"sync"
)
import jsoniter "github.com/json-iterator/go"

// JSON 比原生json性能更好的解析器，用法与原生json一样
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

func NewStudent() *Student {
	return &Student{}
}

type Student struct {
	Name  string
	Age   int
	Right bool
}

func (s *Student) Clear() {
	s.Name = ""
	s.Age = 0
	s.Right = false
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return NewStudent()
	},
}

func main() {
	student := studentPool.Get().(*Student)
	// 使用student

	student.Clear() //返回给studentPool之前必须清空
	studentPool.Put(student)
}
