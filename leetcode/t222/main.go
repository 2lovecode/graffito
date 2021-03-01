// 222. 完全二叉树的节点个数

// 给出一个完全二叉树，求出该树的节点个数。

// 说明：

// 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

// 示例:

// 输入:
//     1
//    / \
//   2   3
//  / \  /
// 4  5 6

// 输出: 6
package main

import "fmt"

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(valSlice []int) *TreeNode {
	var root, newNode, nowNode *TreeNode

	for _, v := range valSlice {
		newNode = &TreeNode{
			Val: v,
		}
		if root == nil {
			root = newNode
		} else {
			nowNode = find(root)
			if nowNode != nil {
				if nowNode.Left == nil {
					nowNode.Left = newNode
				} else if nowNode.Right == nil {
					nowNode.Right = newNode
				}
			}
		}
	}
	return root
}

func find(root *TreeNode) *TreeNode {
	var nowNode *TreeNode

	nowNode = root
	if nowNode == nil {
		return nil
	}

	if nowNode.Left == nil {
		return nowNode
	}

	if nowNode.Right == nil {
		return nowNode
	}

	lNode := find(nowNode.Left)
	if lNode != nil {
		return lNode
	}

	rNode := find(nowNode.Right)
	if rNode != nil {
		return rNode
	}
	return nil
}

func main() {
	fmt.Println(countNodes(NewTree([]int{1, 2, 3, 4, 5, 6})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 1 + countNodes(root.Left) + countNodes(root.Right)

	return count
}
