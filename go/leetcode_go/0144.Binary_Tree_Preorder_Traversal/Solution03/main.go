package main

import (
	"fmt"
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

type stackElem struct {
	node *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	stack := make([]stackElem, 0)
	stack = append(stack, stackElem{node: root})
	for len(stack) > 0 {
		elem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, elem.node.Val)

		if elem.node.Right != nil {
			stack = append(stack, stackElem{node: elem.node.Right})
		}

		if elem.node.Left != nil {
			stack = append(stack, stackElem{node: elem.node.Left})
		}
	}

	return result
}

func test01() {
	root := []int{1, 2, 3}
	successResult := []int{1, 2, 3}
	rootTree := &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", preorderTraversal(rootTree.root))
	fmt.Printf("\n\n")
}

func test02() {
	root := []int{}
	successResult := []int{}
	rootTree := &BinaryTree{}

	fmt.Println("test02: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", preorderTraversal(rootTree.root))
	fmt.Printf("\n\n")
}

func test03() {
	root := []int{1}
	successResult := []int{1}
	rootTree := &BinaryTree{}

	fmt.Println("test03: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", preorderTraversal(rootTree.root))
	fmt.Printf("\n\n")
}

func test04() {
	root := []int{1, 2}
	successResult := []int{1, 2}
	rootTree := &BinaryTree{}

	fmt.Println("test04: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", preorderTraversal(rootTree.root))
	fmt.Printf("\n\n")
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
