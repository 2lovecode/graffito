package apiflow

import (
	"context"
	"fmt"
	"testing"
)

func TestApiFlow_Do(t *testing.T) {
	ctx := context.Background()

	flow := NewApiFlow()

	n1 := NewNode("node1")
	n1.WithHandler(func(ctx context.Context) {
		fmt.Println("1")
	})

	n2 := NewNode("node2")
	n2.WithHandler(func(ctx context.Context) {
		fmt.Println("2")
	})
	n2.WithUpstreams(n1)

	flow.Add(n1)
	flow.Add(n2)

	flow.Start(ctx)
}
