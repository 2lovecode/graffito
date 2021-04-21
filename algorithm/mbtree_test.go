package algorithm

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_MBTree(t *testing.T) {
	Convey("MB-Tree", t, func() {
		size := 100
		var insertP []int

		for _, v := range rand.Perm(size) {
			insertP = append(insertP, v)
		}

		tr := NewMBTree()

		for _, item := range insertP {
			tr.Insert(item)
		}

		fmt.Println(tr.Range(20, 30))
	})
}
