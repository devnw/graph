package graph

import (
	"context"
	"fmt"
	"sync"

	"github.com/benjivesterby/validator"
	"github.com/pkg/errors"
)

// Graphy is the graphy struct for building and searching the graph
type Graphy struct {
	Directional bool
	Weighted    bool

	nodes sync.Map
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
func (g *Graphy) Node(value interface{}) (node Node, err error) {

	if validator.IsValid(value) {
		n := &nodey{
			value: value,
		}

		v, _ := g.nodes.LoadOrStore(value, n)

		var ok bool
		if node, ok = v.(Node); !ok {
			err = errors.Errorf("unable to assert node type for value [%v]", v)
		}
	} else {
		err = errors.Errorf("value [%v] is invalid and not added to graph", value)
	}

	return node, err
}

// RemoveNode removes a node from the graph and removes all edges that reference that node
func (g *Graphy) RemoveNode(value interface{}) (err error) {
	// TODO: validate inputs

	return err
}

// AddEdge adds a new edge to the graph between two nodes
func (g *Graphy) AddEdge(parent Node, child Node, value interface{}, weight int) (err error) {

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
func (g *Graphy) UpdateEdge(parent Node, child Node, value interface{}, weight int) (err error) {

	return err
}

// RemoveEdge removes an edge from the graph
func (g *Graphy) RemoveEdge(parent Node, child Node) (err error) {

	return err
}

func (g *Graphy) String(ctx context.Context) string {
	var output = ""

	g.nodes.Range(func(key, value interface{}) bool {

		if n, ok := value.(Node); ok {
			output = fmt.Sprintf("%s%s\n", output, n.String(ctx))
		}

		// Always loop to completion
		return true
	})

	return output
}
