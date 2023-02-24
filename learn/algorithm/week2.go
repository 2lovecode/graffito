package algorithm

import (
	"github.com/spf13/cobra"
	"graffito/algorithm/link"
)

// 如何判断两个链表相交的位置

func NewIntersectLink() *cobra.Command {
	return &cobra.Command{
		Use: "intersect-link",
		Run: func(cmd *cobra.Command, args []string) {

			// 构造两个链表
			s1 := []int{1, 2, 3, 9}
			s2 := []int{3, 7}

			s3 := []int{9, 10, 8, 10}

			sl1 := link.NewSingleLink[int]()
			sl2 := link.NewSingleLink[int]()
			sl3 := link.NewSingleLink[int]()

			for _, each := range s1 {
				sl1.Append(link.NewNode(each))
			}
			for _, each := range s2 {
				sl2.Append(link.NewNode(each))
			}
			for _, each := range s3 {
				sl3.Append(link.NewNode(each))
			}

		},
	}
}
