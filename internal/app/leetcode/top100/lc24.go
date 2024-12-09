package top100

// 另一种解法 递归法
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	current := prev.Next
	next := current.Next

	for current != nil && next != nil {
		tmp := next.Next
		prev.Next = next
		next.Next = current
		current.Next = tmp

		prev = current
		current = tmp
		if tmp != nil {
			next = tmp.Next
		} else {
			next = nil
		}
	}

	return dummy.Next
}
