package practice

import (
	"github.com/smartystreets/goconvey/convey"
	"graffito/practice/bingfa"
	"testing"
)


func TestBingfa_Run(t *testing.T) {
	convey.Convey("Bingfa_1", t, func() {
		convey.So(bingfa.Run1(), convey.ShouldEqual, 2)
	})
}

func Benchmark_BingfaRun(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		convey.So(bingfa.Run1(), convey.ShouldEqual, 2)
	})
}
