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
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	cur, next := head, head.Next
	for next != nil && next.Val != val {
		cur, next = next, next.Next
	}

	if cur != nil {
		cur.Next = next.Next
	}

	return head
}

func print(nodes *ListNode) {
	for nodes != nil {
		fmt.Print(nodes.Val, "->")
		nodes = nodes.Next
	}
	fmt.Println()
}

func test01() {
	var head = new(ListNode)
	nums := []int{4, 5, 1, 9}
	del := 5
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	print(head)

	head = deleteNode(head, del)

	print(head)
}

func test02() {
	var head = new(ListNode)
	nums := []int{4, 5, 1, 9}
	del := 1
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	print(head)

	head = deleteNode(head, del)

	print(head)
}

func test03() {
	var head = new(ListNode)
	nums := []int{4, 5, 1, 9}
	del := 9
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	print(head)

	head = deleteNode(head, del)

	print(head)
}

func test04() {
	var head = new(ListNode)
	nums := []int{-3, 5, -99}
	del := -3
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	print(head)

	head = deleteNode(head, del)

	print(head)
}

func test05() {
	var head = new(ListNode)
	nums := []int{-5}
	del := -5
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	print(head)

	head = deleteNode(head, del)

	print(head)
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
