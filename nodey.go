package graph

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/benjivesterby/validator"

	"github.com/pkg/errors"
)

// TODO: add the ability to use heuristics through function literals
// TODO: Add support for topographical order
// TODO: Add path loading for reachable nodes
// TODO: Add shortest path loading for reachable nodes

// Node is the interface representation of a node in the graph
// type Node interface {
// 	Edges(ctx context.Context) <-chan Edge
// 	AllReachable(ctx context.Context, alg int) <-chan Node
// 	Reachable(ctx context.Context, alg int, node Node) bool
// 	AddEdge(relation Node, edge Edge) error
// 	DirectMutual(node Node) bool
// 	Value() interface{}
// 	Cost() float64
// 	Parent() Node
// 	String(ctx context.Context) string
// }

// Node is the interface representation of a node in the graph
type Node struct {
	Value  interface{}
	Parent *Node
	Cost   float64
	edges  sync.Map
}

// AllReachable returns all of the nodes that are accessible from this node
// the algorithm passed in is based off of the algorithms defined in the
// CONST files for search constants. Defaults to BFS.
func (n *Node) AllReachable(ctx context.Context, alg int) <-chan *Node {
	var nodes = make(chan *Node)

	go func(nodes chan<- *Node) {
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
func (n *Node) Reachable(ctx context.Context, alg int, node *Node) (reachable bool) {

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
func (n *Node) DirectMutual(node Node) (mutual bool) {

	if validator.IsValid(node) {
		if edge, ok := n.edges.Load(node); ok {
			mutual = validator.IsValid(edge)
		}
	}

	return mutual
}

// AddEdge adds an edge from this node to the related node
func (n *Node) AddEdge(relation *Node, edge Edge) (err error) {
	if validator.IsValid(relation) {
		if validator.IsValid(edge) {

			// TODO: if the edge is weighted should the edge be placed differently in the map?
			if _, loaded := n.edges.LoadOrStore(relation, edge); loaded {
				err = errors.Errorf("edge already exists on node [%v]", n.Value)
			}
		} else {
			err = errors.Errorf("invalid edge [%v]", edge)
		}
	} else {
		err = errors.Errorf("invalid node [%v]", relation)
	}

	return err
}

// Edges returns a channel which streams the edges for this node
func (n *Node) Edges(ctx context.Context) <-chan Edge {
	var edges = make(chan Edge)

	go func(edges chan<- Edge) {
		defer close(edges)

		// Iterate over the edges of the node
		n.edges.Range(func(key, value interface{}) bool {

			select {
			case <-ctx.Done():
				return false
			default:
				if edge, ok := value.(Edge); ok {
					edges <- edge
				}
			}

			// Always loop to completion
			return true
		})

	}(edges)

	return edges
}

func (n *Node) String(ctx context.Context) string {
	var output = "%v: %s"
	var weighted = "(%v, %v)"
	var strs []string

	edges := n.Edges(ctx)

	func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-edges:
				if ok {

					relation := e.Child()
					if relation == n {
						relation = e.Parent()
					}

					switch e.(type) {
					case WeightedEdge:
						if e, ok := e.(WeightedEdge); ok {
							strs = append(strs, fmt.Sprintf(weighted, relation.Value, e.Weight()))
						}
					default:
						strs = append(strs, fmt.Sprintf("%v", relation.Value))
					}
				} else {
					return
				}
			}
		}
	}()

	return fmt.Sprintf(output, n.Value, strings.Join(strs, ","))
}

func (n *Node) Export(ctx context.Context) string {
	var output = "%v=%v"
	var weighted = "%v=%v"
	var strs []string

	edges := n.Edges(ctx)

	func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-edges:
				if ok {

					relation := e.Child()
					if relation == n {
						relation = e.Parent()
					}

					switch e.(type) {
					case WeightedEdge:
						if e, ok := e.(WeightedEdge); ok {
							strs = append(strs, fmt.Sprintf(weighted, relation.Value, e.Weight()))
						}
					default:
						strs = append(strs, fmt.Sprintf("%v", relation.Value))
					}
				} else {
					return
				}
			}
		}
	}()

	return fmt.Sprintf(output, n.Value, strings.Join(strs, "\n"))
}
