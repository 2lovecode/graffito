package top100

type LRUCache struct {
	cap      int
	length   int
	linkHead *linkNode
	linkTail *linkNode
	nodeMap  map[int]*linkNode
}

type linkNode struct {
	key   int
	value int
	prev  *linkNode
	next  *linkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:      capacity,
		length:   0,
		linkHead: nil,
		linkTail: nil,
		nodeMap:  make(map[int]*linkNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		prev := node.prev
		next := node.next

		if prev == nil {
			return node.value
		}

		prev.next = next

		if next != nil {
			next.prev = prev
		} else {
			this.linkTail = prev
		}

		node.prev = nil
		node.next = this.linkHead
		node.next.prev = node
		this.linkHead = node

		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.nodeMap[key]; ok {
		node.value = value
		prev := node.prev
		if prev == nil {
			return
		}
		next := node.next

		prev.next = next

		if next != nil {
			next.prev = prev
		} else {
			this.linkTail = prev
		}

		node.prev = nil
		node.next = this.linkHead
		node.next.prev = node
		this.linkHead = node

	} else {
		if this.length >= this.cap {
			delete(this.nodeMap, this.linkTail.key)
			this.linkTail = this.linkTail.prev
			if this.linkTail != nil {
				this.linkTail.next = nil
			}
			this.length--
		}

		node := &linkNode{
			key:   key,
			value: value,
			prev:  nil,
			next:  this.linkHead,
		}
		this.linkHead = node
		if this.linkTail == nil {
			this.linkTail = node
		}
		if node.next != nil {
			node.next.prev = node
		}
		this.nodeMap[key] = node

		this.length++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
