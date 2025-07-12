package main

import (
	"test26/another"
)

func main() {
	inner := another.NewInner()
	inner.Name()
}

//func main() {
//	errMsg := "123"
//	fmt.Println(errors.Is(errors.New(errMsg), errors.New(errMsg))) //结果反直觉，结果为false
//
//	err123 := errors.New(errMsg)
//	fmt.Println(errors.Is(err123, err123)) //结果为true
//}
