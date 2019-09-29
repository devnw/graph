package graph

// Graphy is the graphy struct for building and searching the graph
type Graphy struct {
	Directional bool
	Weighted    bool

	nodes int
	edges int
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

	g.nodes++

	return node, err
}

// AddEdge adds a new edge to the graph between two nodes
func (g *Graphy) AddEdge(parent Node, child Node, value interface{}, weight int) (Edge, error) {
	var err error

	// TODO: validate inputs

	// TODO: Track most connected node

	var edge = &edgy{
		parent: parent,
		child:  child,
		value:  value,
		weight: weight,
	}

	g.edges++

	return edge, err
}
