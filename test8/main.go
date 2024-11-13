package main

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

type Collection[T any, U any] struct {
	anySlice []T
}

func (c *Collection[T, U]) Get() []T {
	return c.anySlice
}

func NewCollection[T any, U any](array []T) *Collection[T, U] {
	return &Collection[T, U]{
		anySlice: array,
	}
}

func (c *Collection[T, U]) transfer(transFunc func(in T) U) *Collection[U, T] {
	ts := c.Get()
	resArray := make([]U, 0, len(ts))
	for _, t := range ts {
		resArray = append(resArray, transFunc(t))
	}
	return NewCollection[U, T](resArray)
}

func (c *Collection[T, U]) filter(filterFunc func(in T) bool) *Collection[T, U] {
	ts := c.Get()
	resArray := make([]T, 0, len(ts))
	for _, t := range ts {
		if filterFunc(t) {
			resArray = append(resArray, t)

		}
	}
	return NewCollection[T, U](resArray)
}

type Strcut1 struct {
	number string
}

type Strcut2 struct {
	number int
}

func main() {
	struct1Slice := make([]Strcut1, 0)
	for i := 0; i < 100; i++ {
		struct1Slice = append(struct1Slice, Strcut1{cast.ToString(i)})
	}

	res := NewCollection[Strcut1, Strcut2](struct1Slice).transfer(func(t Strcut1) Strcut2 {
		return Strcut2{number: cast.ToInt(t.number)}
	}).filter(func(in Strcut2) bool {
		return in.number >= 50
	}).Get()
	fmt.Println(len(res))
	fmt.Println(res)
}

// mainInner 展示了如果不想使用临时变量应该如何进行类型转变
func mainInner() {
	struct1Slice := make([]Strcut1, 0)
	for i := 0; i < 100; i++ {
		struct1Slice = append(struct1Slice, Strcut1{cast.ToString(i)})
	}

	strcut3Slice := myTransfer(myTransfer(struct1Slice, func(t Strcut1) Strcut2 {
		return Strcut2{number: cast.ToInt(t.number)}
	}), func(t Strcut2) Strcut1 {
		return Strcut1{
			number: cast.ToString(t.number),
		}
	})

	fmt.Println(strcut3Slice)

}

// mainBefore 展示了最原始会反复产生变量的用法，会反复产生局部变量，导致命名，使用等非常的不方便
// 不方便之处：1. 在现在有的代码中存在这么一个情况，a b是两种不同的类型，将aslice转化成bslice，之后，在遍历bslice的过程中有时候用到aslice的变量
// 有时候会用到b.slice   而现有的库pie、collection不能完成对不同结构体的转换。因此经常需要手写
// 2. 变量需要重复的命名。毕竟编程只有两个难题：1.命名；2.缓存失效；3.差一错误
// 3.有些变量可能只是作为中转的变量，命名和查看上下文何处调用都是一个问题
func mainBefore() {
	struct1Slice := make([]Strcut1, 0)
	for i := 0; i < 100; i++ {
		struct1Slice = append(struct1Slice, Strcut1{cast.ToString(i)})
	}
	strcut2Slice := myTransfer(struct1Slice, func(t Strcut1) Strcut2 {
		return Strcut2{number: cast.ToInt(t.number)}
	})

	strcut3Slice := myTransfer(strcut2Slice, func(t Strcut2) Strcut1 {
		return Strcut1{
			number: cast.ToString(t.number),
		}
	})

	fmt.Println(strcut3Slice)

}

func myTransfer[T any, U any](nums []T, trans func(t T) U) []U {
	var zero T
	res := make([]U, 0)
	switch reflect.TypeOf(zero).Kind() {
	case reflect.Struct:
		for _, oneT := range nums {
			oneU := trans(oneT)
			res = append(res, oneU)
		}

	}
	return res
}
