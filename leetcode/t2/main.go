// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

package main

import "fmt"

// ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func main() {
	t1 := []int{9, 9, 9, 9, 9, 9, 9}
	t2 := []int{9, 9, 9, 9}
	var l1, l2, cur *ListNode

	for _, v := range t1 {
		if l1 == nil {
			cur = newNode(v)
			l1 = cur
		} else {
			cur.Next = newNode(v)
			cur = cur.Next
		}
	}

	for _, v := range t2 {
		if l2 == nil {
			cur = newNode(v)
			l2 = cur
		} else {
			cur.Next = newNode(v)
			cur = cur.Next
		}
	}

	printLink(l1)
	printLink(l2)
	printLink(addTwoNumbers(l1, l2))
}

func printLink(x *ListNode) {
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
	fmt.Println("**********")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, cur *ListNode
	jinWei := 0

	for l1 != nil || l2 != nil {
		s := jinWei
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}

		if s < 10 {
			jinWei = 0
		} else {
			jinWei = s / 10
			s = s % 10
		}

		node := &ListNode{
			Val: s,
		}
		if sum == nil {
			cur = node
			sum = cur
		} else {
			cur.Next = node
			cur = cur.Next
		}
	}
	if jinWei != 0 {
		cur.Next = &ListNode{
			Val: jinWei,
		}
		cur = cur.Next
	}
	return sum
}
