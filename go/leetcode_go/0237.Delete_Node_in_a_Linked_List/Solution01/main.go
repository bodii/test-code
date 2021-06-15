package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(n *ListNode) {
	for n != nil {
		fmt.Print(n.Val, "->")
		n = n.Next
	}
	fmt.Println()
}

func deleteNode(node *ListNode) {
	*node = *node.Next
}

func addNode(head []int) *ListNode {
	headNode := new(ListNode)
	headNode.Val = head[0]
	var tailNode *ListNode
	tailNode = headNode

	for i := 1; i < len(head); i++ {
		node := &ListNode{Val: head[i]}
		(*tailNode).Next = node
		tailNode = node
	}

	return headNode
}

func test01() {
	head := []int{4, 5, 1, 9}
	// node := newNode(5)
	// success := []int{4, 1, 9}
	headNode := addNode(head)

	printNode(headNode)
	printNode(headNode)
	deleteNode(headNode)
	printNode(headNode)
}

func main() {
	test01()
}
