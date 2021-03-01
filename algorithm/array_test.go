package algorithm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMyArray(t *testing.T) {
	Convey("创建数组", t, func() {
		myA := NewMyArray(2)
		So(myA, ShouldHaveSameTypeAs, &MyArray{})
		testData := []int{3, 4, 1, 6}
		Convey("Insert&&Get&&Del方法", func() {
			for k, v := range testData {
				myA.Insert(v, k)
			}
			So(myA.Len, ShouldEqual, len(testData))
			So(myA.Get(2), ShouldEqual, 1)

			myA.Del(2)

			So(myA.Len, ShouldEqual, len(testData)-1)
			So(myA.Get(2), ShouldEqual, 6)

		})
	})
}
