// main.go
package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var minCen = 0
var minCenTotalVal = 0

func minimumLevel(root *TreeNode) int {
	list1 := list.New()
	minCen = 1
	minCenTotalVal = root.Val
	list1.PushBack(root)

	curCen := 0
	curCenTotalValue := 0
	for {
		if list1.Len() == 0 {
			break
		}
		curCen++
		curCenTotalValue = 0
		var count = list1.Len()
		for i := 0; i < count; i++ {
			node := list1.Front().Value.(*TreeNode)
			list1.Remove(list1.Front())

			curCenTotalValue += node.Val

			if node.Left != nil {
				list1.PushBack(node.Left)
			}
			if node.Right != nil {
				list1.PushBack(node.Right)
			}
		}
		if curCenTotalValue < minCenTotalVal {
			minCenTotalVal = curCenTotalValue
			minCen = curCen
		}

	}
	return minCen
}
