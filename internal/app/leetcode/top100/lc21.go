package top100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = &ListNode{
				Val: list1.Val,
			}
			current = current.Next

			list1 = list1.Next
		} else if list1.Val == list2.Val {
			current.Next = &ListNode{
				Val: list1.Val,
			}
			current = current.Next

			list1 = list1.Next
			current.Next = &ListNode{
				Val: list2.Val,
			}
			current = current.Next
			list2 = list2.Next
		} else {
			current.Next = &ListNode{
				Val: list2.Val,
			}
			current = current.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return head.Next
}
