package algorithm

import (
	"fmt"
	"github.com/2lovecode/graffito/pkg/algorithm/link"
	"github.com/2lovecode/graffito/pkg/algorithm/tree"
	"github.com/spf13/cobra"
)

// 二叉树中序遍历

func NewBinaryTreeMidOrder() *cobra.Command {
	return &cobra.Command{
		Use: "binary-tree-mid-order",
		Run: func(cmd *cobra.Command, args []string) {
			preList := []int{1, 2, 4, 7, 3, 5, 6, 8}
			midList := []int{4, 7, 2, 1, 5, 3, 8, 6}

			binaryTree := tree.NewBinary[int](preList, midList)
			actualList := midOrder(binaryTree.Root(), len(preList))
			expectList := []int{4, 7, 2, 1, 5, 3, 8, 6}

			fmt.Println("actual: ", actualList)
			fmt.Println("expect: ", expectList)
		},
	}
}

// 二叉树前序遍历

func NewBinaryTreePreOrder() *cobra.Command {
	return &cobra.Command{
		Use: "binary-tree-pre-order",
		Run: func(cmd *cobra.Command, args []string) {
			preList := []int{1, 2, 4, 7, 3, 5, 6, 8}
			midList := []int{4, 7, 2, 1, 5, 3, 8, 6}

			binaryTree := tree.NewBinary[int](preList, midList)
			actualList := preOrder(binaryTree.Root(), len(preList))
			expectList := []int{1, 2, 4, 7, 3, 5, 6, 8}

			fmt.Println("actual: ", actualList)
			fmt.Println("expect: ", expectList)
		},
	}
}

// 二叉树后序遍历

func NewBinaryTreePostOrder() *cobra.Command {
	return &cobra.Command{
		Use: "binary-tree-post-order",
		Run: func(cmd *cobra.Command, args []string) {
			preList := []int{1, 2, 4, 7, 3, 5, 6, 8}
			midList := []int{4, 7, 2, 1, 5, 3, 8, 6}

			binaryTree := tree.NewBinary[int](preList, midList)
			actualList := postOrder(binaryTree.Root())
			expectList := []int{7, 4, 2, 5, 8, 6, 3, 1}

			fmt.Println("actual: ", actualList)
			fmt.Println("expect: ", expectList)
		},
	}
}

func preOrder(root *tree.BinaryNode[int], l int) []int {
	stack := link.NewStack[*tree.BinaryNode[int]](uint(l))
	res := make([]int, 0)

	for root != nil || !stack.Empty() {
		for root != nil {
			_ = stack.Push(root)
			res = append(res, root.Data)
			root = root.Left
		}
		root, _ = stack.Pop()
		root = root.Right
	}
	return res
}

func midOrder(root *tree.BinaryNode[int], l int) []int {
	stack := link.NewStack[*tree.BinaryNode[int]](uint(l))
	res := make([]int, 0)
	for root != nil || !stack.Empty() {
		for root != nil {
			_ = stack.Push(root)
			root = root.Left
		}

		root, _ = stack.Pop()

		res = append(res, root.Data)

		root = root.Right
	}
	return res
}

func postOrder(root *tree.BinaryNode[int]) []int {
	if root == nil {
		return make([]int, 0)
	}
	left := make([]int, 0)
	if root.Left != nil {
		left = postOrder(root.Left)
	}

	right := make([]int, 0)
	if root.Right != nil {
		right = postOrder(root.Right)
	}
	right = append(right, root.Data)
	return append(left, right...)
}
