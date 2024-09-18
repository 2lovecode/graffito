package practice

import (
	"testing"

	"github.com/2lovecode/graffito/internal/app/practice/bingfa"
	"github.com/smartystreets/goconvey/convey"
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
