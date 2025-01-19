package main

import (
	"log"
)

func main() {
	hsize := uintptr(7)
	maxAlign := 16
	x := uintptr(hsize) + uintptr(-int(hsize)&(maxAlign-1))
	log.Print(x)

	log.Print(uintptr(-int(hsize) & (maxAlign - 1)))
	log.Print(-int(hsize) & (maxAlign - 1))
}
