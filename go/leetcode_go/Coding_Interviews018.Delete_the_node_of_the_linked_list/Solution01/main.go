package main

import "fmt"

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
	var root *ListNode
	for head != nil {
		if head.Next != nil && head.Next.Val != val {
			*head.Next = *head.Next.Next
		}

		if head.Next != nil {
			fmt.Println(head.Val, head.Next.Val)
			root.Val = head.Val
			root.Next = head.Next
		}
	}

	return root
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

func main() {
	test01()
	test02()
}
