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
func reverseList(head *ListNode) *ListNode {
	var root *ListNode
	for head != nil {
		head.Next, root, head = root, head, head.Next
	}

	return root
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	resultLen := len(result)
	for i := 0; i < resultLen/2; i++ {
		result[i], result[resultLen-1-i] = result[resultLen-1-i], result[i]
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
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

func test01() {
	var head = new(ListNode)
	nums := []int{1, 2, 3, 4, 5}
	nodes := head
	insert(nodes, nums)
	nodes = head
	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := reversePrint(nodes)
	fmt.Println("test01 success result: [5,4,3,2,1], result:=", result)
	fmt.Println()
}

func test02() {
	var head = new(ListNode)
	nums := []int{1, 2}
	nodes := head
	insert(nodes, nums)
	nodes = head
	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := reversePrint(nodes)
	fmt.Println("test02 success result: [2,1], result:=", result)
	fmt.Println()
}

func test03() {
	var head *ListNode
	nums := []int{}
	nodes := head
	insert(nodes, nums)
	nodes = head
	print(nodes)

	fmt.Println()
	result := reversePrint(nodes)
	fmt.Println("test03 success result: [], result:=", result)
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
