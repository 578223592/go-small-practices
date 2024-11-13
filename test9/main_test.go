package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 定义一个函数，将任意结构体 a（已转换为 any）赋值给结构体 b
func assignStruct(aAsAny any, b any) error {
	aValue := reflect.ValueOf(aAsAny)
	bValue := reflect.ValueOf(b)

	// 检查 b 是否为指针类型，因为我们需要通过指针赋值给它
	if bValue.Kind() != reflect.Ptr {
		return fmt.Errorf("b must be a pointer")
	}

	// 获取 b 指向的值
	bValue = bValue.Elem()

	// 检查 a 和 b 的底层类型是否相同
	if aValue.Type() != bValue.Type() {
		return fmt.Errorf("type mismatch: a (%v) and b (%v) are not the same type", aValue.Type(), bValue.Type())
	}

	// 通过反射将 a 的值赋值给 b
	bValue.Set(aValue)

	return nil
}
func Test_copyStructs(t *testing.T) {

	// 示例：使用不同的结构体
	type Person struct {
		Name string
		Age  int
	}

	type Car struct {
		Brand string
		Year  int
	}

	// 初始化 a 和 b
	a := Person{Name: "Alice", Age: 30}
	var b Person

	// 将 a 转换为 any，并通过 assignStruct 函数赋值给 b
	err := assignStruct(a, &b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("b: %+v\n", b)
	}

	// 示例：Car 结构体
	c := Car{Brand: "Toyota", Year: 2020}
	var d Car

	err = assignStruct(c, &d)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("d: %+v\n", d)
	}

}

func Test2(t *testing.T) {

	// 原始切片
	source := []int{1, 2, 3, 4, 5}
	// 将切片转换为 any
	var a any = source

	// 目标切片，初始化为一个零值切片
	var destination []int

	// 使用反射将 a 赋值给 destination
	// 获取 a 的反射值
	reflectValueOfA := reflect.ValueOf(a)

	// 检查 a 的类型是否为 []int
	if reflectValueOfA.Kind() == reflect.Slice && reflectValueOfA.Type().Elem().Kind() == reflect.Int {
		// 创建目标切片的反射值
		reflectValueOfDest := reflect.MakeSlice(reflect.TypeOf(destination), 0, reflectValueOfA.Len())

		// 逐个复制元素
		for i := 0; i < reflectValueOfA.Len(); i++ {
			reflectValueOfDest = reflect.Append(reflectValueOfDest, reflectValueOfA.Index(i))
		}

		// 将反射值赋值回目标切片
		destination = reflectValueOfDest.Interface().([]int)
	} else {
		fmt.Println("Type mismatch: a is not a []int")
	}

	// 打印结果
	fmt.Println("Destination:", destination)
}
