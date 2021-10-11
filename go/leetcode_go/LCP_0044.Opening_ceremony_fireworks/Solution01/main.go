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
func numColor(root *TreeNode) int {
	res := make(map[int]bool)

	dfs(root, &res)

	return len(res)
}

func dfs(root *TreeNode, res *map[int]bool) {
	if root == nil {
		return
	}

	(*res)[root.Val] = true
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func test01() {
	root := []int{1, 3, 2, 1, 2}
	successResult := 3
	rootTree := &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree length success result:", successResult)
	fmt.Println("result:", numColor(rootTree.root))
	fmt.Println("tree length:", rootTree.count)
	fmt.Println()
}

func test02() {
	root := []int{3, 3, 3}
	successResult := 1
	rootTree := &BinaryTree{}

	fmt.Println("test02: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree length success result:", successResult)
	fmt.Println("result:", numColor(rootTree.root))
	fmt.Println("tree length:", rootTree.count)
	fmt.Println()
}

func test03() {
	root := []int{1, 3, 2, 4, 5, 6}
	successResult := 6
	rootTree := &BinaryTree{}

	fmt.Println("test03: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree length success result:", successResult)
	fmt.Println("result:", numColor(rootTree.root))
	fmt.Println("tree length:", rootTree.count)
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
