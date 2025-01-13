package top100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	//
	length := 23

	step := 1

	for step < length {
	}

	return nil
}

func sort1(head *ListNode, tail *ListNode) *ListNode {
	return nil
}

func sortMerge(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur = list1
			cur = cur.Next
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			cur = list2
			cur = cur.Next
			list2 = list2.Next
		} else {
			cur = list1
			cur = cur.Next
			cur = list2
			cur = cur.Next
			list1 = list1.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		cur = list1
	}
	if list2 != nil {
		cur = list2
	}
	return head
}
