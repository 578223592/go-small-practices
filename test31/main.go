package main

import (
	"log"
)

func main() {
	log.Print(string(1))
	log.Print(string(rune(1)))
	log.Print(string(2))
	log.Print(string(rune(2)))

}

//func main() {
//	errMsg := "123"
//	fmt.Println(errors.Is(errors.New(errMsg), errors.New(errMsg))) //结果反直觉，结果为false
//
//	err123 := errors.New(errMsg)
//	fmt.Println(errors.Is(err123, err123)) //结果为true
//}
