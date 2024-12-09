package top100

// 另外两种解题思路
// 1. 栈：压栈后，弹出的第n个就是要删除的节点，此时栈顶元素即为前驱节点
// 2. 双指针：first先走n个，second再走。当first到尾部时，second正好在倒数第n个节点处

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	m := make(map[int]*ListNode)

	current := head
	cnt := 0
	for current != nil {
		m[cnt] = current
		current = current.Next
		cnt++
	}
	idx := cnt - n
	if idx < 0 {
		return nil
	}
	if idx == 0 {
		return head.Next
	}
	prev := idx - 1
	prevN := m[prev]
	currentN := m[idx]

	prevN.Next = currentN.Next
	return head
}
