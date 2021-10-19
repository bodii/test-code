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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	slow, fast := newHead, newHead
	for ; n >= 0; n-- {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return newHead.Next
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
	n := 2
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := removeNthFromEnd(nodes, n)
	print(result)
	fmt.Println("test01 success result: [1,2,3,5]")
	fmt.Println()
}

func test02() {
	head := new(ListNode)
	l1 := []int{1}
	n := 1
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := removeNthFromEnd(nodes, n)
	print(result)
	fmt.Println("test02 success result: []")
	fmt.Println()
}

func test03() {
	head := new(ListNode)
	l1 := []int{1, 2}
	n := 1
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := removeNthFromEnd(nodes, n)
	print(result)
	fmt.Println("test03 success result: [1]")
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
