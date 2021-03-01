package main

import (
	"context"
	"fmt"
	"graffito/pattern/lib"
)

type Chain1Handler struct {
	lib.Next
}

func (h *Chain1Handler) Do(ctx context.Context) error {
	fmt.Println("i am chain1")
	return nil
}

type Chain2Handler struct {
	lib.Next
}

func (h *Chain2Handler) Do(ctx context.Context) error {
	fmt.Println("i am chain2")
	return nil
}

func main() {
	nilH := &lib.NilHandler{}
	nilH.SetNext(&Chain1Handler{}).SetNext(&Chain2Handler{})
	nilH.Run(context.Background())
}
