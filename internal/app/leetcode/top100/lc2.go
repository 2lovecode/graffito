package top100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	before := 0
	for l1 != nil && l2 != nil {

		sum := l1.Val + l2.Val + before

		val := sum % 10
		before = sum / 10
		current.Next = &ListNode{Val: val}
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next

	}
	for l1 != nil {
		sum := l1.Val + before
		val := sum % 10
		before = sum / 10
		current.Next = &ListNode{Val: val}
		current = current.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + before
		val := sum % 10
		before = sum / 10
		current.Next = &ListNode{Val: val}
		current = current.Next
		l2 = l2.Next
	}
	if before > 0 {
		current.Next = &ListNode{Val: before}
	}

	return head.Next
}
