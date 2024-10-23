package top100

// 234. 回文链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}
	}

	sh := slow

	var pre *ListNode
	for sh != nil {
		next := sh.Next
		sh.Next = pre
		pre = sh
		sh = next
	}

	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}

	return true
}
