package algorithm

import (
	"github.com/spf13/cobra"
)

// 二叉树翻转
// 二叉树最大深度
// 平衡二叉树-判断
// 对称二叉树-判断

func NewReverseBinaryTree() *cobra.Command {
	return &cobra.Command{
		Use: "reverse-binary-tree",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func NewBinaryTreeMaxDepth() *cobra.Command {
	return &cobra.Command{
		Use: "binary-tree-max-depth",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func NewJudgeBalanceBinaryTree() *cobra.Command {
	return &cobra.Command{
		Use: "judge-balance-binary-tree",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func NewJudgeSymmetricalBinaryTree() *cobra.Command {
	return &cobra.Command{
		Use: "judge-symmetrical-binary-tree",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
