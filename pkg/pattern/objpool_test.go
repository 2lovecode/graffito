package pattern

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 对象池模式
func Test_ObjPool(t *testing.T) {
	Convey("ObjPool", t, func() {
		pool := NewPool(10)
		obj := pool.Get()

		obj2 := pool.Get()

		So(obj, ShouldNotBeNil)
		So(obj2, ShouldNotBeNil)
		So(obj.ID(), ShouldNotEqual, obj2.ID())
	})

}
