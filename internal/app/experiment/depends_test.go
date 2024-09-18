package experiment

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/2lovecode/graffito/internal/app/experiment/depends"
)

func TestDepends_Execute(t *testing.T) {
	for i := 0; i < 4; i++ {
		ctx := context.WithValue(context.TODO(), "q", "test")

		sA := depends.NewServiceA()
		sB := depends.NewServiceB()
		sC := depends.NewServiceC()

		hd := depends.NewDepends(100 * time.Millisecond)

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

		fmt.Println(sAData, sBData, sCData)
	}
}

func BenchmarkDepends_Execute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.WithValue(context.TODO(), "q", "test")

		sA := depends.NewServiceA()
		sB := depends.NewServiceB()
		sC := depends.NewServiceC()

		hd := depends.NewDepends(100 * time.Millisecond)

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

		fmt.Println(sAData, sBData, sCData)
	}
}
