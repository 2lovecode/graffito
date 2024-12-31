package apiflow

import (
	"context"
	"fmt"
	"sync"
)

// Handler is a function associated with a node type
// It will be executed during traversal or as needed
// Returns an error to indicate failure
type Handler func(ctx context.Context, node *Node, inputs map[string]interface{}) (interface{}, error)

// Node represents a node in the dependency tree
// Each node corresponds to an API request
// If a node fails, all its dependents are marked as failed

type Node struct {
	ID           string           // Unique identifier of the node
	Predecessors map[string]*Node // Nodes that are dependencies of this node
	Successors   map[string]*Node // Nodes that depend on this node
	Handler      Handler          // Handler function associated with the node
	State        string           // "pending", "success", "failure"
	Data         interface{}      // Data produced by this node
	Mutex        sync.Mutex       // To protect state and data modifications
}

// DependencyTree represents the entire tree structure
type DependencyTree struct {
	nodes map[string]*Node // All nodes in the tree, indexed by their ID
}

// NewDependencyTree creates a new, empty dependency tree
func NewDependencyTree() *DependencyTree {
	return &DependencyTree{
		nodes: make(map[string]*Node),
	}
}

// NewNode creates a new node without adding it to the tree
func NewNode(id string, handler Handler) *Node {
	return &Node{
		ID:           id,
		Predecessors: make(map[string]*Node),
		Successors:   make(map[string]*Node),
		Handler:      handler,
		State:        "pending",
	}
}

// AddNode adds a pre-created node to the dependency tree with its upstream dependencies
func (dt *DependencyTree) AddNode(node *Node, upstreamIDs []string) {
	if _, exists := dt.nodes[node.ID]; exists {
		fmt.Printf("Node %s already exists.\n", node.ID)
		return
	}
	dt.nodes[node.ID] = node

	for _, upstreamID := range upstreamIDs {
		upstreamNode, exists := dt.nodes[upstreamID]
		if !exists {
			fmt.Printf("Upstream node %s does not exist.\n", upstreamID)
			continue
		}
		upstreamNode.Successors[node.ID] = node
		node.Predecessors[upstreamID] = upstreamNode
	}
}
