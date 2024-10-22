package top100

// 206. 反转链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre
}
