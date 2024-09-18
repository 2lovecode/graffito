package experiment

import (
	"testing"

	"github.com/2lovecode/graffito/internal/app/experiment/event"
	"github.com/2lovecode/graffito/pkg/params"

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

		mEvent.Trigger("car", p)
		mEvent.Trigger("phone", p)
	})

}
