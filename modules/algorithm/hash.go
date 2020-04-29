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
	for  k, _ := range d {
		v := h.Get(k)
		if v != "" {
			fmt.Printf("%s, %s\n", k, v)
		}
	}
	fmt.Println("")
}

//开放寻址
type KeyValue struct {
	Key string
	Value string
}

type OpenAddrHash struct{
	List []KeyValue
}

func NewOpenAddrHash(cap int) *OpenAddrHash {
	return &OpenAddrHash{List:make([]KeyValue, cap)}
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

func (h *OpenAddrHash) Get(key string) string{
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
	Key string
	Value string
	Next *LinkNode
}

func NewLinkNode() *LinkNode{
	return &LinkNode{
		Key:   "",
		Value: "",
		Next:  nil,
	}
}

func (node *LinkNode) Find(key string) *LinkNode{
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
	lh := &LinkHash{List:make([]*LinkNode, cap)}
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

func hashFunc(key string, mapLen int) int {
	m := md5.New()
	m.Write([]byte(key))
	sum := 0
	for _, v := range m.Sum([]byte{}) {
		sum += int(v)
	}
	return sum % mapLen
}


func HashRun() {
	/*
		1.数组 -- 可通过下标随机访问 -- O(1)
		2.hash函数计算出的值必须随机且均匀分布 -- 信息论【信息的压缩】
		3.hash函数的时间和空间复杂度
		4.装载因子 -- 衡量冲突的概率 -- 通过扩容重散列降低装载因子
	 */
	testData := map[string]string{
		"k-1" : "v-1",
		"k-2" : "v-2",
		"k-3" : "v-3",
		"k-4" : "v-4",
		"k-5" : "v-5",
		//"k-6" : "v-6",
		//"k-7" : "v-7",
	}
	/**
		开放寻址法解决冲突
		开放寻址，有不同实现
			线性探测 - 实现
			二次探测
			双重散列
	 */
	fmt.Println("===开放寻址处理冲突===")
	openAddrHash := NewOpenAddrHash(5)
	BatchPut(openAddrHash, testData)
	BatchGet(openAddrHash, testData)

	//分离链表处理冲突
	fmt.Println("===分离链表处理冲突===")
	linkHash := NewLinkHash(2)
	BatchPut(linkHash, testData)
	BatchGet(linkHash, testData)


	//装载因子-重散列
	fmt.Println("===分离链表-重散列===")
	linkHash1 := NewLinkHash(2)
	BatchPut(linkHash1, testData)
	BatchGet(linkHash1, testData)

	//
}
