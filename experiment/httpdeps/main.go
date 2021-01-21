package main

import (
	"context"
	"fmt"
	"graffito/experiment/httpdeps/depends"
	"time"
)

func main() {

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

	fmt.Println(sAData, sBData, sCData)

}
