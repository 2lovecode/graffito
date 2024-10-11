package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	convey.Convey("寻找异位词", t, func() {
		convey.So(findAnagrams("cbaebabacd", "abc"), convey.ShouldResemble, []int{0, 6})
	})
}
