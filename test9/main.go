package main

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

var (
	TypeNotSupported = fmt.Errorf(" type is not supported")
	GetValueError    = fmt.Errorf("get value error")
)

type Collection struct {
	InnerSlice          any
	ReflectValueOfInner reflect.Value
	Error               error
}

func (c *Collection) Get(dest any) {
	if c.Error != nil || c.ReflectValueOfInner.Len() <= 0 {
		return
	}
	reflectValueOfDest := reflect.ValueOf(dest)
	if reflectValueOfDest.Kind() != reflect.Ptr || reflectValueOfDest.Elem().Kind() != reflect.Slice {
		c.Error = GetValueError
		return
	}
	reflectValueOfDest = reflectValueOfDest.Elem()
	//if c.ReflectValueOfInner.Type() != reflect.ValueOf(reflectValueOfDest).Type() {
	//	c.Error = GetValueError
	//	return
	//}

	// 确保b的容量足够
	if reflectValueOfDest.Cap() < c.ReflectValueOfInner.Len() {
		reflectValueOfDest.Set(reflect.MakeSlice(reflectValueOfDest.Type(), 0, c.ReflectValueOfInner.Len()))
	}
	reflectValueOfDest.SetLen(0)

	// 逐个复制元素
	for i := 0; i < c.ReflectValueOfInner.Len(); i++ {
		reflectValueOfDest = reflect.Append(reflectValueOfDest, c.ReflectValueOfInner.Index(i))
	}

	//设置回去：因为append过程中reflectValueOfDest有更新，不会自动映射到原来的切片
	reflect.ValueOf(dest).Elem().Set(reflectValueOfDest)
}

func NewCollection(array any) *Collection {
	var err error = nil
	//if len(array) > 0 {
	//	reflectValue := reflect.ValueOf(array[0])
	//	switch reflectValue.Kind() {
	//	case reflect.Slice, reflect.Array: //todo add more type
	//		err = TypeNotSupported
	//	default:
	//
	//	}
	//
	//}
	return &Collection{
		InnerSlice:          array,
		ReflectValueOfInner: reflect.ValueOf(array),
		Error:               err,
	}
}

// transfer
//
//	@Description: 对保存的元素逐个进行transFunc
//	@receiver c
//	@param transFunc
//	@return *Collection
func (c *Collection) transfer(transFunc func(in any) (out any)) *Collection {
	if c.Error != nil || c.ReflectValueOfInner.Len() <= 0 {
		return c
	}
	oldReflectValueOfInner := c.ReflectValueOfInner
	//oldInnerSlice := c.InnerSlice
	//outType, err := GetReturnType(transFunc) 失败的方案1:因为这个func就是返回的any，所以虽然用了反射，但是依然获取不到类型
	//
	//if err != nil {
	//	c.Error = err
	//	return c
	//}
	outType := reflect.TypeOf(transFunc(oldReflectValueOfInner.Index(0).Interface()))
	fmt.Printf("transfer Output Type:%s\n", outType.Name())

	c.ReflectValueOfInner = reflect.MakeSlice(reflect.SliceOf(outType), oldReflectValueOfInner.Len(), oldReflectValueOfInner.Len())
	//c.InnerSlice = make([]any, reflect.ValueOf(c.InnerSlice).Len(), reflect.ValueOf(c.InnerSlice).Len()) //这里是不是应该将any转成out 的类型
	//newReflectValue := reflect.ValueOf(c.InnerSlice)
	for i := 0; i < c.ReflectValueOfInner.Len(); i++ {
		newElem := transFunc(oldReflectValueOfInner.Index(i).Interface())
		valueOf := reflect.ValueOf(newElem)
		//fmt.Println(reflect.TypeOf(newElem).Name())
		c.ReflectValueOfInner.Index(i).Set(valueOf)
	}
	return c
}

// GetReturnType 获取函数返回值类型的函数
func GetReturnType(fn func(in any) (out any)) (reflect.Type, error) {
	// 获取函数的反射类型
	funcType := reflect.TypeOf(fn)

	// 检查是否是函数类型
	if funcType.Kind() != reflect.Func {
		fmt.Println("Input is not a function")
		return nil, fmt.Errorf("input is not a function")
	}

	// 获取函数的第一个返回值类型
	// 函数可以有多个返回值，所以我们用 Out(0) 获取第一个返回值
	return0Name := funcType.Out(0).Name()
	fmt.Println(return0Name)
	return funcType.Out(0), nil
}

//func (c *Collection[T, U]) filter(filterFunc func(in T) bool) *Collection[T, U] {
//	ts := c.Get()
//	resArray := make([]T, 0, len(ts))
//	for _, t := range ts {
//		if filterFunc(t) {
//			resArray = append(resArray, t)
//
//		}
//	}
//	return NewCollection[T, U](resArray)
//}

type Strcut1 struct {
	numberOfStruct1 string
	name            string
}

type Strcut2 struct {
	numberOfStruct2 int
}

func main() {
	struct1Slice := make([]Strcut1, 0)
	for i := 0; i < 100; i++ {
		struct1Slice = append(struct1Slice, Strcut1{numberOfStruct1: cast.ToString(i)})
	}
	//strcut2Slice := make([]Strcut2, 0)

	NewCollection(struct1Slice).transfer(func(in any) any {
		strcut2 := Strcut2{numberOfStruct2: cast.ToInt(in.(Strcut1).numberOfStruct1) + 10}
		return strcut2
	}).transfer(func(in any) any {
		strcut1 := Strcut1{numberOfStruct1: cast.ToString(in.(Strcut2).numberOfStruct2)}
		return strcut1
	}).Get(&struct1Slice)

	fmt.Println(len(struct1Slice))
	fmt.Println(struct1Slice)
}

//
//func main2() {
//	struct1Slice := make([]Strcut1, 0)
//	for i := 0; i < 100; i++ {
//		struct1Slice = append(struct1Slice, Strcut1{cast.ToString(i)})
//	}
//	strcut2Slice := make([]Strcut2, 0)
//
//	for _, strcut1 := range struct1Slice {
//		//fliter
//
//		tmpsRuct2 := Strcut2{
//			number: cast.ToInt(strcut1.number),
//		}
//		strcut2Slice = append(strcut2Slice)
//
//	}
//
//	fmt.Println(len(strcut2Slice))
//	fmt.Println(strcut2Slice)
//
//}
