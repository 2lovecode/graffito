// 147. 对链表进行插入排序

// 对链表进行插入排序。

// 插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
// 每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

// 插入排序算法：

//     插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//     每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//     重复直到所有输入数据插入完为止。

// 示例 1：

// 输入: 4->2->1->3
// 输出: 1->2->3->4

// 示例 2：

// 输入: -1->5->3->4->0
// 输出: -1->0->3->4->5
package leetcode

import "fmt"

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func main() {
	l := []int{4, 2, 1, 3}
	var head, newNode, nowNode *ListNode
	for _, v := range l {
		newNode = &ListNode{
			Val: v,
			Next: nil,
		}
		if head == nil {
			head = newNode
			nowNode = newNode
		} else {
			nowNode.Next = newNode
			nowNode = nowNode.Next
		}
	}

	nodePrint(insertionSortList(head)) 
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	var nowNode, sortedNode, prev *ListNode
	if head == nil {
		return head
	}
	nowNode = head.Next
	sortedNode = head

	emptyHead := &ListNode{Next:head}

	for nowNode != nil {
		if sortedNode.Val <= nowNode.Val {
			sortedNode = sortedNode.Next
		} else {
			prev = emptyHead
			for prev.Next.Val <= nowNode.Val {
				prev = prev.Next
			}
			sortedNode.Next = nowNode.Next
			nowNode.Next = prev.Next
			prev.Next = nowNode
		}

		nowNode = sortedNode.Next
	}

	return emptyHead.Next
}

func nodePrint(head *ListNode) {
	var nowNode *ListNode
	nowNode = head
	for nowNode != nil {
		fmt.Println(nowNode.Val)
		nowNode = nowNode.Next
	}
}
