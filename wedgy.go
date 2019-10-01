package graph

// wedgy is the weighted edge
type wedgy struct {
	parent Node
	child  Node
	value  interface{}
	weight int
}

func (e *wedgy) Parent() Node {
	return e.parent
}

func (e *wedgy) Child() Node {
	return e.child
}

func (e *wedgy) Value() interface{} {
	return e.value
}

func (e *wedgy) Weight() int {
	return e.weight
}
