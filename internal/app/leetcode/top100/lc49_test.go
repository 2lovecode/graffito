package top100

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	convey.Convey("异位词", t, func() {
		fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	})
}

func TestGroupAnagrams2(t *testing.T) {
	convey.Convey("异位词2", t, func() {
		fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	})
}
