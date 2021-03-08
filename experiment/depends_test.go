package experiment

import (
	"context"
	"graffito/experiment/depends"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Depends(t *testing.T) {
	Convey("Depends", t, func() {
		ctx := context.WithValue(context.TODO(), "q", "test")

		sA := depends.NewServiceA()
		sB := depends.NewServiceB()
		sC := depends.NewServiceC()

		hd := depends.NewHttpDeps(100 * time.Millisecond)

		hd.Register(sA)
		hd.Register(sB)
		hd.Register(sC)

		hd.AddDepend(sC, []depends.IService{sB})
		hd.AddDepend(sB, []depends.IService{sA})

		hd.Execute(ctx)

		sAData := depends.ServiceAData{}
		sBData := depends.ServiceBData{}
		sCData := depends.ServiceCData{}

		sA.Decode(&sAData)
		sB.Decode(&sBData)
		sC.Decode(&sCData)
		So(sAData.Message, ShouldEqual, "service_a")
		So(sBData.Message, ShouldEqual, "service_b")
		So(sCData.Message, ShouldEqual, "service_c")
	})
}
