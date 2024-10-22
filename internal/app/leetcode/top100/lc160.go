package top100

type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. 相交链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA := headA
	currentB := headB
	for currentA != currentB {
		if currentA == nil {
			currentA = headB
		} else {
			currentA = currentA.Next
		}

		if currentB == nil {
			currentB = headA
		} else {
			currentB = currentB.Next
		}
	}

	return currentA
}
