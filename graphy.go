// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package graph

import (
	"context"
	"fmt"
	"sync"

	"github.com/devnw/validator"
	"github.com/pkg/errors"
)

// Graphy is the graphy struct for building and searching the graph
type Graphy struct {
	Directional bool
	Weighted    bool

	nodes     sync.Map
	size      int
	sizeMutty sync.Mutex

	edges sync.Map
}

// DAG determines if the graph that's build is a Directed Acyclic Graph
func (g *Graphy) DAG() (isDAG bool) {

	return isDAG
}

// Strong determines if the graph has strong connectivity
func (g *Graphy) Strong() (isStrong bool) {

	return isStrong
}

// Bipartite determines if the graph is bipartite
func (g *Graphy) Bipartite() (isBipartite bool) {

	return isBipartite
}

// Mutual accepts two nodes of the graph and determines if they're mutually reachable
func (g *Graphy) Mutual(n1 Node, n2 Node) (mutual bool) {
	// TODO: validate inputs

	return mutual
}

// Node adds a new node to the graph if it
// doesn't already exist and returns the node
// if the node already exists then it returns
// the node object from the map
func (g *Graphy) Node(value interface{}) (node *Node, err error) {

	if validator.IsValid(value) {
		n := &Node{
			Value: value,
		}

		v, loaded := g.nodes.LoadOrStore(value, n)

		// increment size
		if !loaded {
			g.sizeMutty.Lock()
			g.size++
			g.sizeMutty.Unlock()
		}

		var ok bool
		if node, ok = v.(*Node); !ok {
			err = errors.Errorf("unable to assert node type for value [%v]", v)
		}
	} else {
		err = errors.Errorf("value [%v] is invalid and not added to graph", value)
	}

	return node, err
}

func (g *Graphy) AddNode(node *Node) (err error) {
	if validator.IsValid(node) {
		_, loaded := g.nodes.LoadOrStore(node.Value, node)

		// increment size
		if !loaded {
			g.sizeMutty.Lock()
			g.size++
			g.sizeMutty.Unlock()
		}
	} else {
		// TODO:
	}

	return err
}

// Size returns the current size of the graph
func (g *Graphy) Size() int {
	g.sizeMutty.Lock()
	defer g.sizeMutty.Unlock()

	s := g.size
	return s
}

// Nodes returns a full set of the nodes in the graph with their associated edges
func (g *Graphy) Nodes(ctx context.Context) <-chan *Node {
	nodes := make(chan *Node)

	go func(nodes chan<- *Node) {
		defer close(nodes)

		g.nodes.Range(func(key, value interface{}) bool {

			if n, ok := value.(*Node); ok {
				if n != nil {
					// Push the node onto the channel
					select {
					case <-ctx.Done():
						// Break the loop
						return false
					case nodes <- n:
					}
				} else {
					// TODO:
				}
			}

			// Always loop to completion
			return true
		})
	}(nodes)

	return nodes
}

// RemoveNode removes a node from the graph and removes all edges that reference that node
func (g *Graphy) RemoveNode(value interface{}) (err error) {
	// TODO: validate inputs
	g.sizeMutty.Lock()
	defer g.sizeMutty.Unlock()

	g.size--
	// TODO: Implement removal

	return err
}

// AddEdge adds a new edge to the graph between two nodes
func (g *Graphy) AddEdge(parent, child *Node, value interface{}, weight float64) (err error) {

	// TODO: validate inputs

	// TODO: Track most connected node

	var edge Edge
	// Build the edge based on the type of graph
	if g.Directional {
		if g.Weighted {
			edge = &dwedgy{
				parent: parent,
				child:  child,
				value:  value,
				weight: weight,
			}
		} else {
			edge = &dedgy{
				parent: parent,
				child:  child,
				value:  value,
			}
		}
	} else {
		if g.Weighted {
			edge = &wedgy{
				parent: parent,
				child:  child,
				value:  value,
				weight: weight,
			}
		} else {
			edge = &edgy{
				parent: parent,
				child:  child,
				value:  value,
			}
		}

		if parent != child {
			// TODO: Register the edge in the child node
			// TODO: Register the edge in the edge map for the child index
			if err = child.AddEdge(parent, edge); err == nil {

				// TODO: deal with the sync map here. Need to take into account a possible duplicate edge...
				if _, loaded := g.edges.LoadOrStore(child, edge); !loaded {
					// TODO:
				} else {
					// TODO: Error here because the edge already existed
				}
			}
		}
	}

	//TODO: Register the edge in the parent node
	//TODO: Register the edge in the edge map for the parent index
	if err = parent.AddEdge(child, edge); err == nil {

		// TODO: deal with the sync map here. Need to take into account a possible duplicate edge...
		if _, loaded := g.edges.LoadOrStore(parent, edge); !loaded {
			// TODO:
		} else {
			// TODO: Error here because the edge already existed
		}
	}

	return err
}

// UpdateEdge updates the information in an edge for the graph
func (g *Graphy) UpdateEdge(parent *Node, child *Node, value interface{}, weight int) (err error) {

	return err
}

// RemoveEdge removes an edge from the graph
func (g *Graphy) RemoveEdge(parent *Node, child *Node) (err error) {

	return err
}

func (g *Graphy) String(ctx context.Context) string {
	var output = ""

	nodes := g.Nodes(ctx)

	// Setup function literal to break out of when the loop completes
	func() {
		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-nodes:
				if ok {
					output = fmt.Sprintf("%s%s\n", output, n.String(ctx))
				} else {
					return
				}
			}
		}
	}()

	return output
}

func (g *Graphy) Export(ctx context.Context) string {

	direction := "undirected"
	if g.Directional {
		direction = "directed"
	}

	weighted := "unweighted"
	if g.Weighted {
		weighted = "weighted"
	}

	output := fmt.Sprintf("%s %s\n", direction, weighted)

	filter := make(map[interface{}]bool)

	g.nodes.Range(func(key, value interface{}) bool {

		if n, ok := value.(*Node); ok {

			func() {
				edges := n.Edges(ctx)
				for {
					select {
					case <-ctx.Done():
						return
					case e, ok := <-edges:
						if ok {
							if !filter[e.Child().Value] && !filter[e.Parent().Value] {
								edge := fmt.Sprintf("%s%v=%v\n", output, e.Parent().Value, e.Child().Value)
								if we, ok := e.(WeightedEdge); ok {
									edge = fmt.Sprintf("%s%v=%v=%v\n", output, e.Parent().Value, e.Child().Value, we.Weight())
								}

								output = edge
							}
						} else {
							return
						}
					}
				}
			}()

			filter[n.Value] = true
		}

		return true
	})

	// g.edges.Range(func(key, value interface{}) bool {

	// 	if e, ok := value.(Edge); ok {
	// 		if e != nil {
	// 			key := fmt.Sprintf("%v%v", e.Child().Value, e.Parent().Value)
	// 			key2 := fmt.Sprintf("%v%v", e.Parent().Value, e.Child().Value)

	// 			fmt.Println(key, key2, !filter[key] && !filter[key2])

	// 			if !filter[key] && !filter[key2] {

	// 				edge := fmt.Sprintf("%s%v=%v\n", output, e.Parent().Value, e.Child().Value)
	// 				if we, ok := e.(WeightedEdge); ok {
	// 					edge = fmt.Sprintf("%s%v=%v=%v\n", output, e.Parent().Value, e.Child().Value, we.Weight())
	// 				}

	// 				output = edge
	// 				// Update the filter that this edge has been seen
	// 				filter[key] = true
	// 				filter[key2] = true
	// 			}
	// 		} else {
	// 			// TODO:
	// 		}
	// 	}

	// 	// Always loop to completion
	// 	return true
	// })

	// nodes := g.Nodes(ctx)

	// // Setup function literal to break out of when the loop completes
	// func() {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		case n, ok := <-nodes:
	// 			if ok {
	// 				output = fmt.Sprintf("%s%s\n", output, n.Export(ctx))
	// 			} else {
	// 				return
	// 			}
	// 		}
	// 	}
	// }()

	return output
}
