package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = result<<1 + head.Val
		head = head.Next
	}

	return result
}

func insert(nodes *ListNode, nums []int) {
	numsLen := len(nums)
	if numsLen < 1 {
		return
	}

	nodes.Val = nums[0]
	for i := 1; i < numsLen; i++ {
		node := &ListNode{Val: nums[i]}
		nodes.Next = node
		nodes = node
	}
}

func print(nodes *ListNode) {
	for nodes != nil {
		fmt.Println("print node:", nodes.Val)
		nodes = nodes.Next
	}
}

func test01() {
	head := new(ListNode)
	l1 := []int{1, 0, 1}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test01 success result: 5")
	fmt.Println("result: ", getDecimalValue(nodes))
	fmt.Println()
}

func test02() {
	head := new(ListNode)
	l1 := []int{0}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test02 success result: 0")
	fmt.Println("result: ", getDecimalValue(nodes))
	fmt.Println()
}

func test03() {
	head := new(ListNode)
	l1 := []int{1}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test03 success result: 1")
	fmt.Println("result: ", getDecimalValue(nodes))
	fmt.Println()
}

func test04() {
	head := new(ListNode)
	l1 := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test04 success result: 18880")
	fmt.Println("result: ", getDecimalValue(nodes))
	fmt.Println()
}

func test05() {
	head := new(ListNode)
	l1 := []int{0, 0}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test05 success result: 0")
	fmt.Println("result: ", getDecimalValue(nodes))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
