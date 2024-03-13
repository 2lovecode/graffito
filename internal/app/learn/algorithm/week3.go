package algorithm

import (
	"fmt"
	"github.com/2lovecode/graffito/pkg/algorithm/link"
	"github.com/spf13/cobra"
)

// 单链表反转
func NewReverseLink() *cobra.Command {
	return &cobra.Command{
		Use: "reverse-link",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("反转单链表：")
			fmt.Println("step1: 生成单链表")
			var head *link.Node[int]
			head = GenSingleLinkedList([]int{1, 3, 4, 5, 9})

			LinkPrint(head)
			fmt.Println("")
			fmt.Println("step2: 反转后")
			head = LinkReverse(head)
			LinkPrint(head)
		},
	}
}

// 单链表倒数第k个节点
func NewKthFromBottom() *cobra.Command {
	return &cobra.Command{
		Use: "kth-from-bottom",
		Run: func(cmd *cobra.Command, args []string) {
			// 1. 遍历获得链表长度，计算正数节点数，再次遍历
			// 2. 快慢指针，快指针先移动k，随后快慢指针一起移动，快指针到达尾部后，慢指针指向倒数k
			linkedList := GenSingleLinkedList([]int{1, 2, 3, 4, 5, 6, 7})

			var fast *link.Node[int]
			var slow *link.Node[int]

			k := 2

			fast = linkedList
			slow = linkedList

			for i := 0; i < k; i++ {
				fast = fast.Next()
			}
			for fast != nil {
				fast = fast.Next()
				slow = slow.Next()
			}
			fmt.Println(slow.Value())
		},
	}
}
