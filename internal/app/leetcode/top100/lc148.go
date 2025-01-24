package top100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return top2bottomMergeSort(head, nil)
}

func top2bottomMergeSort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	mid := middleListNode(head, tail)

	return mergeTwoSortedList(top2bottomMergeSort(head, mid), top2bottomMergeSort(mid, tail))
}

func bottom2topMergeSort(head *ListNode) *ListNode {
	length := listLength(head)

	dummyHead := &ListNode{Next: head}

	for subLength := 1; subLength < length; subLength *= 2 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			nHead, nTail := mergeTwoSortedList2(head1, head2)

			prev.Next = nHead
			prev = nTail

			cur = next
		}
	}

	return dummyHead.Next
}

func listLength(head *ListNode) int {
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}
	return l
}

func middleListNode(head *ListNode, tail *ListNode) *ListNode {
	fast, slow := head, head

	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	return slow
}

func mergeTwoSortedList(head1 *ListNode, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}

	temp, temp1, temp2 := dummyHead, head1, head2

	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	}

	if temp2 != nil {
		temp.Next = temp2
	}

	return dummyHead.Next
}

func mergeTwoSortedList2(head1 *ListNode, head2 *ListNode) (*ListNode, *ListNode) {
	dummyHead := &ListNode{}

	temp, temp1, temp2 := dummyHead, head1, head2

	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	}

	if temp2 != nil {
		temp.Next = temp2
	}

	var tail *ListNode

	for temp != nil {
		tail = temp
		temp = temp.Next
	}

	return dummyHead.Next, tail
}
