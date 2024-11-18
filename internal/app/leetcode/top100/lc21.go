package top100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := &ListNode{}

	//for list1 != nil && list2 != nil {
	//	if list1.Val < list2.Val {
	//		list.Next = list1
	//		list1 = list1.Next
	//	} else if list1.Val == list2.Val {
	//		list.Next = list1
	//		list1 = list1.Next
	//		list.Next = list2
	//		list2 = list2.Next
	//	} else {
	//		list.Next = list2
	//		list2 = list2.Next
	//	}
	//}
	//if list1 != nil {
	//	list.Next = list1
	//}
	//if list2 != nil {
	//	list.Next = list2
	//}

	return list
}
