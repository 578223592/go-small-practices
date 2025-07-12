package main

import (
	"log"
)

type CustomizedError struct {
}

func (c CustomizedError) Error() string {
	return "CustomizedError"
}

func func1() error {
	var err *CustomizedError = nil
	return err
}

func main() {
	err := func1()
	if err != nil {
		log.Fatalf("1 err: %v", err) // 以为 err 为nil，但是会进这个分支
	}
}

//func main() {
//	errMsg := "123"
//	fmt.Println(errors.Is(errors.New(errMsg), errors.New(errMsg))) //结果反直觉，结果为false
//
//	err123 := errors.New(errMsg)
//	fmt.Println(errors.Is(err123, err123)) //结果为true
//}
