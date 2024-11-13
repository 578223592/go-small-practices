package main

type FloatData interface {
	~float32 | ~float64
}

type IntData interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}

// Number This is used as a type constraint for the generic Tensor type.
type Number interface {
	FloatData | IntData
}
type myStruct1[T Number] struct {
	Data []T
}

// 根据传入的参数，返回包含不同数据类型的myStruct1类型的实例
func getMyStruct1[T Number](choice int) myStruct1[T] {
	if choice == 1 {
		return myStruct1[int64]{Data: {1, 2, 3}}
	} else {
		return myStruct1[int32]{Data: {2, 3, 4}}
	}
}

// 泛型函数不能直接返回不同类型的结构体实例
// 修改为具体类型后再返回
func getMyStruct2(choice int) any {
	if choice == 1 {
		return myStruct1[int64]{Data: []int64{1, 2, 3}}
	} else {
		return myStruct1[int32]{Data: []int32{2, 3, 4}}
	}
}

// 将参数类型提前到泛型参数T上，通过调用不同的泛型实例化函数来获得不同的结果
func getMyStruct3[T Number]() myStruct1[T] {
	// 根据T的类型来返回不同的结果
	var result myStruct1[T]
	switch any(result).(type) {
	case myStruct1[int64]:
		result = myStruct1[T]{Data: []T{1, 2, 3}}
	case myStruct1[int32]:
		result = myStruct1[T]{Data: []T{2, 3, 4}}
	// 如果需要，可以继续增加其他类型的匹配
	default:
		// 默认返回空的myStruct1
		result = myStruct1[T]{Data: []T{}}
	}
	return result
}
