package top100

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	// fmt.Println(cache.linkHead.value, cache.linkTail.value)
	fmt.Println(cache.Get(1))
	// fmt.Println(cache.linkHead.value, cache.linkTail.value)

	cache.Put(3, 3)
	// fmt.Println(cache.linkHead.value, cache.linkTail.value)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
