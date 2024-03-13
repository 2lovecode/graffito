package pattern

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Observer(t *testing.T) {
	Convey("Observer", t, func() {
		notifier := NewNotifier()

		ob1 := NewObserver(1)
		ob2 := NewObserver(2)

		notifier.Register(ob1)
		notifier.Register(ob2)

		actual := 223

		notifier.Notify(Event{Data: actual})

		So(actual, ShouldEqual, ob1.Data)
		So(actual, ShouldEqual, ob2.Data)

	})

}
