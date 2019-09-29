package graph

import (
	"context"
)

// TODO: add the ability to use heuristics through function literals
// TODO: Add support for topographical order

// Node is the interface representation of a node in the graph
type Node interface {
	Edges() <-chan *Edge
}

type nodey struct {
	value interface{}
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

// Reachable determines if a node is reachable from this node
func (n *nodey) Reachable(ctx context.Context, alg int, node Node) (reachable bool) {

	var nodes = n.AllReachable(ctx, alg)

	for {
		select {
		case <-ctx.Done():
			return
		case n, ok := <-nodes:
			if ok {
				if n == node {
					reachable = true
					return
				}
			} else {
				return
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

func (n *nodey) Edges(ctx context.Context) <-chan Edge {
	var edges = make(chan Edge)

	go func(edges chan<- Edge) {
		defer close(edges)

	}(edges)

	return edges
}
