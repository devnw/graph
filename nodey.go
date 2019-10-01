package graph

import (
	"context"
	"sync"

	"github.com/pkg/errors"
)

// TODO: add the ability to use heuristics through function literals
// TODO: Add support for topographical order

// Node is the interface representation of a node in the graph
type Node interface {
	Edges() <-chan Edge
	AllReachable(ctx context.Context, alg int) <-chan Node
	Reachable(ctx context.Context, alg int, node Node) bool
	AddEdge(relation Node, edge Edge) error
	DirectMutual(node Node) bool
	Value() interface{}
}

type nodey struct {
	value interface{}
	edges sync.Map
}

// AllReachable returns all of the nodes that are accessible from this node
// the algorithm passed in is based off of the algorithms defined in the
// CONST files for search constants. Defaults to BFS.
func (n *nodey) AllReachable(ctx context.Context, alg int) <-chan Node {
	var nodes = make(chan Node)

	go func(nodes chan<- Node) {
		defer close(nodes)

		switch alg {
		case DFS:
			// TODO: Call out to depth first search for results
		case BFS:
			fallthrough
		default:
			// TODO: Call out to breadth first search
		}

	}(nodes)

	return nodes
}

// TODO: Add path to the return from reachable. This will need to be implemented using the parallel DFS and BFS options

// Reachable determines if a node is reachable from this node
func (n *nodey) Reachable(ctx context.Context, alg int, node Node) (reachable bool) {

	var nodes = n.AllReachable(ctx, alg)

	for {
		select {
		case <-ctx.Done():
			return reachable
		case n, ok := <-nodes:
			if ok {
				if n == node {
					reachable = true
					return reachable
				}
			} else {
				return reachable
			}
		}
	}
}

// DirectMutual determines if a specific node is directly mutual to this node
func (n *nodey) DirectMutual(node Node) (mutual bool) {
	return mutual
}

func (n *nodey) Value() interface{} {
	return n.value
}

func (n *nodey) AddEdge(relation Node, edge Edge) (err error) {

	// TODO: if the edge is weighted should the edge be placed differently in the map?
	if _, loaded := n.edges.LoadOrStore(relation, edge); loaded {
		err = errors.Errorf("edge already exists on node %v", n.value)
	}

	return err
}

// Edges returns a channel which streams the edges for this node
func (n *nodey) Edges(ctx context.Context) <-chan Edge {
	var edges = make(chan Edge)

	go func(edges chan<- Edge) {
		defer close(edges)

		// Iterate over the edges of the node
		n.edges.Range(func(key, value interface{}) bool {

			select {
			case <-ctx.Done():
				return false
			default:
				if value != nil {
					if edge, ok := value.(Edge); ok {
						edges <- edge
					}
					// ignore else statement here on purpose. If the value is not an edge just leave it
				} else {
					// delete this record from the map because it's nil
					n.edges.Delete(key)
				}
			}

			// Always loop to completion
			return true
		})

	}(edges)

	return edges
}
