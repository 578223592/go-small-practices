package main

import "fmt"

type Node struct {
	val int
}

func (receiver Node) changeVal1() {
	receiver.val++
}

func (receiver *Node) changeVal2() {
	receiver.val++
}
func main() {
	someOne := Node{}
	someOne.changeVal1()
	fmt.Printf("%d\n", someOne.val)

	someOnePtr := &Node{}
	someOnePtr.changeVal1()
	fmt.Printf("%d\n", someOnePtr.val)

	someTwo := Node{}
	someTwo.changeVal2()
	fmt.Printf("%d\n", someTwo.val)

	someTwoPtr := &Node{}
	someTwoPtr.changeVal2()
	fmt.Printf("%d\n", someTwoPtr.val)
}
