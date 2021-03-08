package algorithm

import (
	"crypto/md5"
	"fmt"
)

type HashInterface interface {
	Put(key string, value string)
	Get(key string) string
}

func BatchPut(h HashInterface, d map[string]string) {
	for k, v := range d {
		h.Put(k, v)
	}
}

func BatchGet(h HashInterface, d map[string]string) {
	for k, _ := range d {
		v := h.Get(k)
		if v != "" {
			fmt.Printf("%s, %s\n", k, v)
		}
	}
	fmt.Println("")
}

//开放寻址
type KeyValue struct {
	Key   string
	Value string
}

type OpenAddrHash struct {
	List []KeyValue
}

func NewOpenAddrHash(cap int) *OpenAddrHash {
	return &OpenAddrHash{List: make([]KeyValue, cap)}
}

func (h *OpenAddrHash) Put(key string, value string) {
	hashLen := len(h.List)
	hashKey := hashFunc(key, hashLen)
	originHashKey := hashKey

	for h.List[hashKey] != (KeyValue{}) {
		hashKey = (hashKey + 1) % hashLen
		if hashKey == originHashKey {
			break
		}
	}
	if h.List[hashKey] != (KeyValue{}) {
		fmt.Println("hash is full!")
	} else {
		h.List[hashKey] = KeyValue{
			Key:   key,
			Value: value,
		}
	}
}

func (h *OpenAddrHash) Get(key string) string {
	hashLen := len(h.List)
	hashKey := hashFunc(key, hashLen)
	originHashKey := hashKey

	for h.List[hashKey].Key != key {
		hashKey = (hashKey + 1) % hashLen
		if hashKey == originHashKey {
			break
		}
	}
	if h.List[hashKey].Key != key {
		return ""
	} else {
		return h.List[hashKey].Value
	}
}

//分离链表法
type LinkNode struct {
	Key   string
	Value string
	Next  *LinkNode
}

func NewLinkNode() *LinkNode {
	return &LinkNode{
		Key:   "",
		Value: "",
		Next:  nil,
	}
}

func (node *LinkNode) Find(key string) *LinkNode {
	current := node
	for current.Next != nil {
		if current.Key == key {
			break
		} else {
			current = current.Next
		}
	}
	return current
}

func (node *LinkNode) Add(key string, value string) {
	current := node.Find(key)
	if current.Key == key {
		current.Value = value
	} else {
		newLinkNode := NewLinkNode()
		newLinkNode.Key = key
		newLinkNode.Value = value
		current.Next = newLinkNode
	}
}

type LinkHash struct {
	List []*LinkNode
}

func NewLinkHash(cap int) *LinkHash {
	lh := &LinkHash{List: make([]*LinkNode, cap)}
	for i := 0; i < cap; i++ {
		lh.List[i] = NewLinkNode()
	}
	return lh
}

func (lh *LinkHash) Put(key string, value string) {
	mapLen := len(lh.List)
	hashKey := hashFunc(key, mapLen)
	lh.List[hashKey].Find(key).Add(key, value)
}

func (lh *LinkHash) Get(key string) string {
	mapLen := len(lh.List)
	hashKey := hashFunc(key, mapLen)
	return lh.List[hashKey].Find(key).Value
}

type LinkReHash struct {
	List []*LinkNode
	Len  int
}

func NewLinkReHash(cap int) *LinkReHash {
	lh := &LinkReHash{
		List: make([]*LinkNode, cap),
		Len:  0,
	}
	for i := 0; i < cap; i++ {
		lh.List[i] = NewLinkNode()
	}
	return lh
}

func (lh *LinkReHash) Put(key string, value string) {
	mapLen := len(lh.List)
	if float32(lh.Len)/float32(mapLen) >= 1.2 {
		newMapLen := mapLen * 2
		newList := make([]*LinkNode, newMapLen)
		for i := 0; i < newMapLen; i++ {
			newList[i] = NewLinkNode()
		}
		for _, v := range lh.List {
			for v.Next != nil {
				v = v.Next
				newHashKey := hashFunc(v.Key, newMapLen)
				newList[newHashKey].Find(v.Key).Add(v.Key, v.Value)
			}
		}
		lh.List = newList
		mapLen = newMapLen
	}
	hashKey := hashFunc(key, mapLen)
	lh.List[hashKey].Find(key).Add(key, value)
	lh.Len++
}

func (lh *LinkReHash) Get(key string) string {
	mapLen := len(lh.List)
	hashKey := hashFunc(key, mapLen)
	return lh.List[hashKey].Find(key).Value
}

func hashFunc(key string, mapLen int) int {
	m := md5.New()
	m.Write([]byte(key))
	sum := 0
	for _, v := range m.Sum([]byte{}) {
		sum += int(v)
	}
	return sum % mapLen
}
