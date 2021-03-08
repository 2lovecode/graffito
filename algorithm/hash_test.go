package algorithm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOpenAddrHash(t *testing.T) {
	Convey("创建开放寻址哈希表", t, func() {
		h := NewOpenAddrHash(5)
		So(h, ShouldHaveSameTypeAs, &OpenAddrHash{})
		Convey("Put和Get", func() {
			testData := map[string]string{
				"k-1": "v-1",
				"k-2": "v-2",
				"k-3": "v-3",
				"k-4": "v-4",
				"k-5": "v-5",
				//"k-6" : "v-6",
				//"k-7" : "v-7",
			}
			BatchPut(h, testData)
			for k, v := range testData {
				So(h.Get(k), ShouldEqual, v)
			}
		})
	})
}

func TestLinkHash(t *testing.T) {
	Convey("创建分离链表哈希表", t, func() {
		h := NewLinkHash(5)
		So(h, ShouldHaveSameTypeAs, &LinkHash{})
		Convey("Put和Get", func() {
			testData := map[string]string{
				"k-1": "v-1",
				"k-2": "v-2",
				"k-3": "v-3",
				"k-4": "v-4",
				"k-5": "v-5",
				"k-6": "v-6",
				"k-7": "v-7",
			}
			BatchPut(h, testData)
			for k, v := range testData {
				So(h.Get(k), ShouldEqual, v)
			}
		})
	})
}

func TestLinkReHash(t *testing.T) {
	Convey("创建分离链表重散列哈希表", t, func() {
		h := NewLinkReHash(5)
		So(h, ShouldHaveSameTypeAs, &LinkReHash{})
		Convey("Put和Get", func() {
			testData := map[string]string{
				"k-1":  "v-1",
				"k-2":  "v-2",
				"k-3":  "v-3",
				"k-4":  "v-4",
				"k-5":  "v-5",
				"k-6":  "v-6",
				"k-7":  "v-7",
				"k-8":  "v-8",
				"k-9":  "v-9",
				"k-10": "v-10",
				"k-11": "v-11",
				"k-12": "v-12",
				"k-13": "v-13",
			}
			BatchPut(h, testData)
			for k, v := range testData {
				So(h.Get(k), ShouldEqual, v)
			}
		})
	})
}
