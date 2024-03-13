package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_QuickSort(t *testing.T) {
	Convey("QuickSort", t, func() {
		s := []int{4, 5, 6, 3, 7, 9}
		Quick(s, 0, len(s)-1)
		So(s[0], ShouldEqual, 3)
		So(s[5], ShouldEqual, 9)
	})

}
