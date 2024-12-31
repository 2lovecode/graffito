package apiflow

import (
	"context"
	"errors"
	"fmt"
	"github.com/sourcegraph/conc"
	"time"
)

// ApiFlow wraps a dependency tree to provide API execution logic
type ApiFlow struct {
	tree        *DependencyTree
	resultsChan chan *Node
	timeout     time.Duration
	cancel      context.CancelFunc
	wg          *conc.WaitGroup
}

// NewApiFlow creates a new ApiFlow instance
func NewApiFlow(timeout time.Duration) *ApiFlow {
	return &ApiFlow{
		tree:        NewDependencyTree(),
		resultsChan: make(chan *Node, 100), // Buffered to avoid blocking
		timeout:     timeout,
	}
}

// AddNode adds a node with dependencies to the flow
func (af *ApiFlow) AddNode(node *Node, upstreamIDs []string) {
	af.tree.AddNode(node, upstreamIDs)
}

// Run executes the API flow
func (af *ApiFlow) Run(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, af.timeout)
	af.cancel = cancel
	defer cancel()

	af.wg = conc.NewWaitGroup()

	// Start worker goroutines to process nodes
	af.wg.Go(func() {
		af.processResults(ctx)
	})

	// Traverse the tree and start execution of root nodes
	for _, node := range af.tree.nodes {
		if len(node.Predecessors) == 0 {
			af.wg.Go(func() {
				af.executeNode(ctx, node)
			})
		}
	}

	af.wg.Wait()
	close(af.resultsChan)
}

// executeNode runs the handler for a single node and sends the result to the results channel
func (af *ApiFlow) executeNode(ctx context.Context, node *Node) {

	node.Mutex.Lock()
	if node.State != "pending" {
		node.Mutex.Unlock()
		return
	}
	node.Mutex.Unlock()

	// Collect input data from predecessors
	inputs := make(map[string]interface{})
	for id, predecessor := range node.Predecessors {
		predecessor.Mutex.Lock()
		inputs[id] = predecessor.Data
		predecessor.Mutex.Unlock()
	}

	// Execute the handler with a timeout
	nodeCtx, cancel := context.WithTimeout(ctx, af.timeout)
	defer cancel()

	data, err := node.Handler(nodeCtx, node, inputs)
	node.Mutex.Lock()
	if err != nil || errors.Is(nodeCtx.Err(), context.DeadlineExceeded) {
		if errors.Is(nodeCtx.Err(), context.DeadlineExceeded) {
			fmt.Printf("Node %s timed out\n", node.ID)
		}
		node.State = "failure"
	} else {
		node.State = "success"
		node.Data = data // Store the result of the handler
	}
	node.Mutex.Unlock()

	af.resultsChan <- node
}

// processResults processes completed nodes and triggers successors
func (af *ApiFlow) processResults(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case node, ok := <-af.resultsChan:
			if !ok {
				return
			}
			noPending := true
			for _, n := range af.tree.nodes {
				n.Mutex.Lock()
				if n.State == "pending" {
					noPending = false
				}
				n.Mutex.Unlock()
			}
			if noPending {
				af.cancel()
				return
			}

			if node.State == "success" {
				for _, successor := range node.Successors {
					successorReady := true
					successor.Mutex.Lock()
					for _, predecessor := range successor.Predecessors {
						if predecessor.State != "success" {
							successorReady = false
							break
						}
					}
					successor.Mutex.Unlock()
					if successorReady {
						af.wg.Go(func() {
							af.executeNode(ctx, successor)
						})
					}
				}
			} else if node.State == "failure" {
				for _, successor := range node.Successors {
					successor.Mutex.Lock()
					if successor.State == "pending" {
						successor.State = "failure"
						fmt.Printf("Marking node %s as failed due to dependency failure\n", successor.ID)
					}
					successor.Mutex.Unlock()
					af.resultsChan <- successor
				}
			}
		}
	}
}
