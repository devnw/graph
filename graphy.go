package graph

import "sync"

// Graphy is the graphy struct for building and searching the graph
type Graphy struct {
	Directional bool
	Weighted    bool

	nodeCountLock sync.Mutex
	nodeCount     int
	nodes         sync.Map

	edgeCountLock sync.Mutex
	edgeCount     int
	edges         sync.Map
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

// AddNode adds a new node to the graph using the information passed in
func (g *Graphy) AddNode(value interface{}) (node Node, err error) {
	// TODO: validate inputs

	// increment the total node count for O(1) counting
	g.nodeCountLock.Lock()
	g.nodeCount++
	g.nodeCountLock.Unlock()

	return node, err
}

// RemoveNode removes a node from the graph and removes all edges that reference that node
func (g *Graphy) RemoveNode(value interface{}) (err error) {
	// TODO: validate inputs

	// decrement the total node count for O(1) counting
	g.nodeCountLock.Lock()
	g.nodeCount--
	g.nodeCountLock.Unlock()

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
			edge = &directedWedgy{
				parent: parent,
				child:  child,
				value:  value,
				weight: weight,
			}
		} else {
			edge = &directedEdgy{
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

		// TODO: Register the edge in the child node
		// TODO: Register the edge in the edge map for the child index
	}

	//TODO: Register the edge in the parent node
	//TODO: Register the edge in the edge map for the parent index
	if err = parent.AddEdge(child, edge); err == nil {

		// TODO: deal with the sync map here. Need to take into account a possible duplicate edge...
		if _, loaded := g.edges.LoadOrStore(parent, edge); !loaded {

			// TODO: determine if this should remain here
			// increment the total edge count for O(1) counting
			g.edgeCountLock.Lock()
			g.edgeCount++
			g.edgeCountLock.Unlock()
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

	// decrement the total edge count for O(1) counting
	g.edgeCountLock.Lock()
	g.edgeCount--
	g.edgeCountLock.Unlock()

	return err
}
