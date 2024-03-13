package other

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Fibonacci(t *testing.T) {
	Convey("Fibonacci", t, func() {
		So(Fibonacci(FibonacciCnt), ShouldEqual, 89)

		So(FibonacciUseMem(FibonacciCnt), ShouldEqual, 89)
	})

}
