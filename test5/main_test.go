package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

var bytes, _ = json.Marshal(NewStudent())

func BenchmarkNewWithoutPoolNewWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		student := NewStudent()
		_ = json.Unmarshal(bytes, student)
	}
}

func BenchmarkNewWithoutPoolNewWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		student := studentPool.Get().(*Student)
		_ = json.Unmarshal(bytes, student)
		studentPool.Put(student)
	}
}

// json的marshall测试
func TestPoolNewValue(t *testing.T) {
	studentWithAge := NewStudent()
	studentWithAge.Age = 18
	agebuff, _ := JSON.Marshal(studentWithAge) // {"Name":"","Age":18,"Right":false} 说明对于零值的字段，仍然会序列化
	fmt.Printf(string(agebuff) + "\n")

	student := studentPool.Get().(*Student)
	student.Name = "你好"

	fmt.Printf("unmarshal之前：" + student.Name + strconv.Itoa(student.Age) + "\n")

	_ = JSON.Unmarshal(agebuff, student)
	fmt.Printf("unmarshal之后：" + student.Name + strconv.Itoa(student.Age) + "\n") //因为上面agebuff对零值的字段仍然会序列化，因此这里会打印出"18"，即name为空

	studentPool.Put(student)
	student = studentPool.Get().(*Student)
	fmt.Printf(student.Name + strconv.Itoa(student.Age))
	if student.Name == "你好" {
		t.Fatal("error" + student.Name)
	} else {
		t.Log("ok")
	}
}
