package skiplist

import "github.com/2lovecode/graffito/pkg/trait"

type List[K trait.Comparable, V any] struct {
	start    []*Node[K, V]
	maxLevel int
}

type Node[K trait.Comparable, V any] struct {
	key   K
	value V
	next  []*Node[K, V]
}

func NewList[K trait.Comparable, V any](maxLevel int) *List[K, V] {
	return &List[K, V]{
		maxLevel: maxLevel,
	}
}

func (l *List[K, V]) Insert(key K, value V) {

}

func (l *List[K, V]) Find(key K) *Node[K, V] {

	startLevel := 0
	for i, n := range l.start {
		if n.key.Eq(key) || n.key.Less(key) {
			startLevel = i
			break
		}
	}
	println(startLevel)
	return nil

}

func (l *List[K, V]) Delete(key K) {

}
