package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	convey.Convey("最长连续序列", t, func() {
		convey.So(longestConsecutive([]int{100, 4, 200, 1, 3, 2}), convey.ShouldEqual, 4)
	})
	//fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
