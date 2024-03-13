package link

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack_PushAndPop(t *testing.T) {

	Convey("压栈和出栈", t, func() {
		stack := NewStack[int](5)
		list := []int{1, 2, 3, 4, 5}
		expect := []int{5, 4, 3, 2, 1}

		for _, each := range list {
			pushErr := stack.Push(each)
			So(pushErr, ShouldEqual, nil)
		}
		for _, each := range expect {
			v, popErr := stack.Pop()
			So(popErr, ShouldEqual, nil)
			So(v, ShouldEqual, each)
		}
	})

}
