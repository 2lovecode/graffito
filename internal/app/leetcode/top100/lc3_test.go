package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	convey.Convey("无重复子串长度", t, func() {
		convey.So(lengthOfLongestSubstring("abcabcbb"), convey.ShouldEqual, 3)
		convey.So(lengthOfLongestSubstring("bbbbb"), convey.ShouldEqual, 1)
		convey.So(lengthOfLongestSubstring("pwwkew"), convey.ShouldEqual, 3)
		convey.So(lengthOfLongestSubstring("b"), convey.ShouldEqual, 1)
		convey.So(lengthOfLongestSubstring("au"), convey.ShouldEqual, 2)
		convey.So(lengthOfLongestSubstring("dvdf"), convey.ShouldEqual, 3)
	})

	convey.Convey("无重复子串长度2", t, func() {
		convey.So(lengthOfLongestSubstring2("abcabcbb"), convey.ShouldEqual, 3)
		convey.So(lengthOfLongestSubstring2("bbbbb"), convey.ShouldEqual, 1)
		convey.So(lengthOfLongestSubstring2("pwwkew"), convey.ShouldEqual, 3)
		convey.So(lengthOfLongestSubstring2("b"), convey.ShouldEqual, 1)
		convey.So(lengthOfLongestSubstring2("au"), convey.ShouldEqual, 2)
		convey.So(lengthOfLongestSubstring2("dvdf"), convey.ShouldEqual, 3)
	})
}
