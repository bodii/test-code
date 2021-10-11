package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root  *TreeNode
	count int
}

func (tn *TreeNode) addNode(newNode *TreeNode) {
	if tn.Val > newNode.Val {
		if tn.Left == nil {
			tn.Left = newNode
		} else {
			tn.Left.addNode(newNode)
		}
	} else {
		if tn.Right == nil {
			tn.Right = newNode
		} else {
			tn.Right.addNode(newNode)
		}
	}
}

func (bt *BinaryTree) add(val int) {
	var newNode TreeNode
	newNode.Val = val
	if bt.root == nil {
		bt.root = &newNode
	} else {
		bt.root.addNode(&newNode)
	}
	bt.count++
}

func (tn *TreeNode) printNode() {
	if tn.Left != nil {
		tn.Left.printNode()
		fmt.Print("left val: ", tn.Left.Val)
	}
	fmt.Print(" ->val: ", tn.Val)

	if tn.Right != nil {
		tn.Right.printNode()
		fmt.Print("right val: ", tn.Right.Val)
	}
	fmt.Println()
}

func (bt *BinaryTree) printTree() {
	if bt.root == nil {
		return
	}

	bt.root.printNode()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func test01() {
	root := []int{10, 5, 15, 3, 7, 18}
	low, high := 7, 15
	successResult := 32
	rootTree := &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree sum success result:", successResult)
	fmt.Println("result:", rangeSumBST(rootTree.root, low, high))
	fmt.Println("tree length:", rootTree.count)
	fmt.Println()
}

func test02() {
	root := []int{10, 5, 15, 3, 7, 13, 18, 1, 6}
	low, high := 6, 10
	successResult := 23
	rootTree := &BinaryTree{}

	fmt.Println("test02: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree sum success result:", successResult)
	fmt.Println("result:", rangeSumBST(rootTree.root, low, high))
	fmt.Println("tree length:", rootTree.count)
	fmt.Println()
}

func main() {
	test01()
	test02()
}
