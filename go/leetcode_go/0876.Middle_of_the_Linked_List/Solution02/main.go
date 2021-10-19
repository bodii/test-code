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
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
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
		fmt.Println("print:", nodes.Val)
		nodes = nodes.Next
	}
}

func test01() {
	head := new(ListNode)
	l1 := []int{1, 2, 3, 4, 5}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := middleNode(nodes)
	print(result)
	fmt.Println("test01 success result: [3,4,5]")
	fmt.Println()
}

func test02() {
	head := new(ListNode)
	l1 := []int{1, 2, 3, 4, 5, 6}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := middleNode(nodes)
	print(result)
	fmt.Println("test02 success result: [4,5,6]")
	fmt.Println()
}

func main() {
	test01()
	test02()
}
