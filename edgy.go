package graph

// Edge is the interface that defines an edge in the graph
type Edge interface {
	Parent() Node
	Child() Node
	Directional() bool
	Weight() int
	Value() interface{}
}

// the struct that implements the Edge interface in the graph
// library. It's a bit edgy :P
type edgy struct {
	parent      Node
	child       Node
	directional bool
	weight      int
	value       interface{}
}

func (e *edgy) Value() interface{} {
	return e.value
}

func (e *edgy) Directional() bool {
	return e.directional
}

func (e *edgy) Parent() (node Node) {
	if !e.directional {
		node = e.parent
	}

	return node
}

func (e *edgy) Child() Node {
	return e.child
}

func (e *edgy) Weight() int {
	return e.weight
}
