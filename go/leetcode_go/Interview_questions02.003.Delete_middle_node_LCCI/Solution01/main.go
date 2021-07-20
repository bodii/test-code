package main

import "fmt"

var head = new(ListNode)

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
func deleteNode(node *ListNode) {
	*node = *node.Next
}

func print(nodes *ListNode) {
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

func main() {
	nums := []int{4, 5, 1, 9}
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	print(head)

	nodes = head
	for nodes != nil {
		if nodes.Val == 5 {
			deleteNode(nodes)
		}
		nodes = nodes.Next
	}

	print(head)
}
