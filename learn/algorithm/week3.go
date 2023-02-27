package algorithm

import (
	"fmt"
	"github.com/spf13/cobra"
	"graffito/algorithm/link"
)

// 单链表反转
func NewReverseLink() *cobra.Command {
	return &cobra.Command{
		Use: "reverse-link",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("反转单链表：")
			fmt.Println("step1: 生成单链表")
			var head *link.Node[int]
			var current *link.Node[int]

			lv := []int{1, 3, 4, 5, 9}
			for _, each := range lv {
				if head == nil {
					head = link.NewNode(each)
					current = head
				} else {
					newNode := link.NewNode(each)
					current.SetNext(newNode)
					current = newNode
				}
			}
			linkPrint(head)
			fmt.Println("")
			fmt.Println("step2: 反转后")
			head = linkReverse(head)
			linkPrint(head)
		},
	}
}

func linkPrint(now *link.Node[int]) {
	if now != nil {
		now.Print()
		fmt.Print(",")
		now = now.Next()
		linkPrint(now)
	}
}

func linkReverse(head *link.Node[int]) *link.Node[int] {
	var pre *link.Node[int]
	current := head

	for current != nil {
		next := current.Next()
		current.SetNext(pre)

		pre = current
		current = next
	}
	return pre
}
