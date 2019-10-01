package graph

// Edge is the interface that defines an edge in the graph
type Edge interface {
	Parent() Node
	Child() Node
}

// DirectionalEdge is the interface that defines a directional edge in the graph
type DirectionalEdge interface {
	Edge
	Directional() bool
	Value() interface{}
}

// DirectionalWeightedEdge is the interface that defines a directional weighted edge in the graph
type DirectionalWeightedEdge interface {
	DirectionalEdge
	Weight() int
}

// WeightedEdge is the interface that defines an undirected weighted edge in the graph
type WeightedEdge interface {
	Edge
	Weight() int
}
