package algorithm

import "github.com/spf13/cobra"

// 位运算
// 数组中只出现一次的数（这个数只有一个）
// 0-n，每个数中二进制表示中1的个数
// 计算两个数之间的汉明距离

func NewOnlyOnceInArray() *cobra.Command {
	return &cobra.Command{
		Use: "only-once-in-array",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func NewBitOneNum() *cobra.Command {
	return &cobra.Command{
		Use: "bit-one-num",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func NewCalcHammingDistance() *cobra.Command {
	return &cobra.Command{
		Use: "calc-hamming-distance",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
