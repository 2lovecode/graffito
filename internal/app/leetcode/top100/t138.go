package top100

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var newHead *Node
	newMap := make(map[*Node]*Node)
	if head != nil {
		current := head
		for current != nil {
			newNode := &Node{Val: current.Val, Next: current.Next, Random: current.Random}
			newMap[current] = newNode
			current = current.Next
		}
		for k, v := range newMap {
			if v.Next != nil {
				if _, ok := newMap[v.Next]; ok {
					newMap[k].Next = newMap[v.Next]
				}
			}
			if v.Random != nil {
				if _, ok := newMap[v.Random]; ok {
					newMap[k].Random = newMap[v.Random]
				}
			}
		}
	}

	if head != nil {
		if _, ok := newMap[head]; ok {
			newHead = newMap[head]
		}
	}
	return newHead
}
