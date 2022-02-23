package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Print(" ->val: ", tn.Val)
	if tn.Left != nil {
		tn.Left.printNode()
		fmt.Print(" l: ", tn.Left.Val)
	}

	if tn.Right != nil {
		tn.Right.printNode()
		fmt.Print(" r: ", tn.Right.Val)
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
func tree2str(root *TreeNode) string {
	result := make([]byte, 0)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, []byte(strconv.Itoa(root.Val))...)

		if root.Left != nil {
			result = append(result, '(')
			dfs(root.Left)
			result = append(result, ')')
		}

		if root.Left == nil && root.Right != nil {
			result = append(result, '(', ')')
		}

		if root.Right != nil {
			result = append(result, '(')
			dfs(root.Right)
			result = append(result, ')')
		}
	}

	dfs(root)

	return string(result)
}

func test01() {
	root := []int{1, 2, 3, 4}
	successResult := "1(2(4))(3)"
	rootTree := &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", tree2str(rootTree.root))
	fmt.Printf("\n\n")
}

func test02() {
	root := []int{1, 2, 3, 4}
	successResult := "1(2()(4))(3)"
	rootTree := &BinaryTree{}

	fmt.Println("test02: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", tree2str(rootTree.root))
	fmt.Printf("\n\n")
}

func test03() {
	root := []int{-1, -2, -3, -4}
	successResult := "-1(-2()(-4))(-3)"
	rootTree := &BinaryTree{}

	fmt.Println("test03: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", tree2str(rootTree.root))
	fmt.Printf("\n\n")
}

func test04() {
	root := []int{37, -34, -48, -100, -100, 48, -54, -71, -22, 8}
	successResult := "37(-34()(-100))(-48(-100)(48(-54(-71)(-22()(8)))))"
	rootTree := &BinaryTree{}

	fmt.Println("test04: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", tree2str(rootTree.root))
	fmt.Printf("\n\n")
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
