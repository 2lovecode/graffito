package algorithm


import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func perm(n int) (out []Item) {
	for _, v := range rand.Perm(n) {
		out = append(out, Int(v))
	}
	return
}

func Test_BTree(t *testing.T) {
	Convey("B-Tree", t, func() {
		size := 100
		insertP := perm(size)
		tr := New(32)
		for _, item := range insertP {
			tr.ReplaceOrInsert(item)
		}

		var got []Item


		tr.AscendGreaterOrEqual(Int(40), func(a Item) bool {
			got = append(got, a)
			return true
		})

		fmt.Println(got)
	})

}
