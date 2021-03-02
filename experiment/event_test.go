package experiment

import (
	"graffito/experiment/event"
	"graffito/utils/params"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Event(t *testing.T) {
	Convey("Event", t, func() {
		mEvent := event.NewMEvent()

		myC := &event.MyCar{}
		myP := &event.MyPhone{}

		mEvent.On("car", myC)
		mEvent.On("phone", myP)

		p := params.NewPayload()
		p.Set("name", "vivo")

		So(mEvent.Trigger("car", p), ShouldEqual, "car-vivo")
		So(mEvent.Trigger("phone", p), ShouldEqual, "phone-vivo")
	})

}
