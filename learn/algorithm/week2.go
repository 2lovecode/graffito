package algorithm

import (
	"fmt"
	"github.com/spf13/cobra"
	"graffito/algorithm/link"
)

// 如何判断两个链表相交的位置

func NewIntersectLink() *cobra.Command {
	return &cobra.Command{
		Use: "intersect-link",
		Run: func(cmd *cobra.Command, args []string) {
			a, b := link.GenTwoIntersectLink(8, 13, 5)

			fmt.Println("方法一：暴力破解 O(m*n),O(1)")

			fmt.Println("方法二：hashmap O(m+n),O(n)")

			fmt.Println("方法三：双指针 O(m+n),O(1)")
			p1 := a.Head()
			p2 := b.Head()
			for p1 != p2 {
				if p1 == nil {
					p1 = p2
				}
				if p2 == nil {
					p2 = p1
				}
				p1 = p1.Next()
				p2 = p2.Next()
			}
			if p1 != nil {
				p1.Print()
			}

			fmt.Println("方法四：对齐长度遍历 O(m+n),O(1)")
		},
	}
}
