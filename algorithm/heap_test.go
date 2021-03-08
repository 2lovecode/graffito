package algorithm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Heap(t *testing.T) {
	Convey("最大堆实现", t, func() {
		heap := NewIntHeap(5)
		testData := []int{1, 9, 4, 10, 22, 14, 8, 7}
		for _, v := range testData {
			heap.Push(v)
		}

		So(heap.Pop(), ShouldEqual, 22)
		So(heap.Pop(), ShouldEqual, 14)
	})

}
