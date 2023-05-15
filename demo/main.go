package main

import (
	"errors"
	"fmt"
	"sync"
)

type LRUCache struct {
	lock sync.RWMutex
	Hash map[int]*Node
	List *LinkedList
}

func NewLRUCache() *LRUCache {
	return &LRUCache{
		lock: sync.RWMutex{},
		Hash: make(map[int]*Node),
		List: NewLinkedList(),
	}
}

func (lc *LRUCache) Put(key int, value int) (err error) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	node, ok := lc.Hash[key]

	if ok && node != nil {
		// 在当前位置移动到最后
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}

		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		lc.List.Tail.Next = node
		node.Prev = lc.List.Tail
		node.Next = nil
	} else {
		if lc.List.Len >= lc.List.Cap {
			// 删除头部节点
			lc.List.Head = lc.List.Head.Next
			lc.List.Head.Prev = nil
			lc.List.Len--
		}

		newNode := &Node{
			Key:   key,
			Value: value,
			Prev:  lc.List.Tail,
			Next:  nil,
		}

		if lc.List.Head == nil {
			lc.List.Head = newNode
			lc.List.Tail = newNode
		} else {
			lc.List.Tail.Next = newNode

			lc.Hash[key] = newNode
		}

	}
	return
}

func (lc *LRUCache) Get(key int) (value int, err error) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	node, ok := lc.Hash[key]
	// 获取到则将节点移到tail
	if ok && node != nil {
		value = node.Value

		// 在当前位置移动到最后
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}

		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		lc.List.Tail.Next = node
		node.Prev = lc.List.Tail
		node.Next = nil

	} else {
		err = errors.New("no data")
	}

	return
}

type LinkedList struct {
	Len  int
	Cap  int
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Len:  0,
		Cap:  3,
		Head: nil,
		Tail: nil,
	}
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

// lru cache get/put
func main() {
	testKeyData := []int{2, 4, 6, 8}
	testValueData := []int{100, 200, 300, 400}
	cache := NewLRUCache()

	// test put
	for k, v := range testKeyData {
		if len(testValueData) > k {
			fmt.Println("error: ", cache.Put(v, testValueData[k]))
		}
	}

	// test get
	expect := 200
	actual, _ := cache.Get(4)
	fmt.Println("expect value: ", expect, "actual value: ", actual)
}
