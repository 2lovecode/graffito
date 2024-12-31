package apiflow

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestApiFlow_Do(t *testing.T) {
	ctx := context.Background()

	flow := NewApiFlow(10 * time.Second)

	nodeA := NewNode("A", func(ctx context.Context, node *Node, inputs map[string]interface{}) (interface{}, error) {
		fmt.Printf("Executing node %s\n", node.ID)
		time.Sleep(1 * time.Second)
		m := make(map[string]string)
		m["node_a_1"] = "goodA"
		m["node_a_2"] = "morningA"
		return m, nil
	})

	nodeB := NewNode("B", func(ctx context.Context, node *Node, inputs map[string]interface{}) (interface{}, error) {
		fmt.Printf("Executing node %s\n", node.ID)
		if inputs != nil {
			if aa, ok := inputs["A"]; ok {
				fmt.Println("bbbbbb: ", aa)
			}
		}
		time.Sleep(2 * time.Second)
		m := make(map[string]string)
		m["node_b_1"] = "goodB"
		m["node_b_2"] = "morningB"
		return m, nil
	})

	nodeC := NewNode("C", func(ctx context.Context, node *Node, inputs map[string]interface{}) (interface{}, error) {
		fmt.Printf("Executing node %s\n", node.ID)
		if inputs != nil {
			if aa, ok := inputs["A"]; ok {
				fmt.Println("cccc: ", aa)
			}
			if bb, ok := inputs["B"]; ok {
				fmt.Println("cccc: ", bb)
			}
		}
		time.Sleep(1 * time.Second)
		return nil, fmt.Errorf("error in node %s", node.ID)
	})

	nodeD := NewNode("D", func(ctx context.Context, node *Node, inputs map[string]interface{}) (interface{}, error) {
		fmt.Printf("Executing node %s\n", node.ID)
		time.Sleep(3 * time.Second)
		return nil, nil
	})

	// Adding nodes with dependencies
	flow.AddNode(nodeA, []string{})
	flow.AddNode(nodeB, []string{"A"})
	flow.AddNode(nodeC, []string{"A"})
	flow.AddNode(nodeD, []string{"A"})

	// Run the API flow
	flow.Run(ctx)
}
