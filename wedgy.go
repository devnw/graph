package graph

// wedgy is the weighted edge
type wedgy struct {
	parent *Node
	child  *Node
	value  interface{}
	weight float64
}

func (e *wedgy) Parent() *Node {
	return e.parent
}

func (e *wedgy) Child() *Node {
	return e.child
}

func (e *wedgy) Value() interface{} {
	return e.value
}

func (e *wedgy) Weight() float64 {
	return e.weight
}
