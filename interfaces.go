package graph

// Edge is the interface that defines an edge in the graph
type Edge interface {
	Parent() Node
	Child() Node
}

// DirectedEdge is the interface that defines a directional edge in the graph
type DirectedEdge interface {
	Edge
	Directional() bool
	Value() interface{}
}

// DirectedWeightedEdge is the interface that defines a directional weighted edge in the graph
type DirectedWeightedEdge interface {
	DirectedEdge
	Weight() int
}

// WeightedEdge is the interface that defines an undirected weighted edge in the graph
type WeightedEdge interface {
	Edge
	Weight() int
}
