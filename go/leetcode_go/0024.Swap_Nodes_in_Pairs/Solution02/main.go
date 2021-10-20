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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head.Next
	head.Next = swapPairs(cur.Next)
	cur.Next = head
	return cur
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
	l1 := []int{1, 2, 3, 4}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := swapPairs(nodes)
	print(result)
	fmt.Println("test01 success result: [2,1,4,3]")
	fmt.Println()
}

func test02() {
	head := new(ListNode)
	l1 := []int{}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := swapPairs(nodes)
	print(result)
	fmt.Println("test02 success result: []")
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
	result := swapPairs(nodes)
	print(result)
	fmt.Println("test03 success result: [1]")
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
