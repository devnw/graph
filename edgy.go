package graph

type edgy struct {
	parent *Node
	child  *Node
	value  interface{}
}

func (e *edgy) Parent() *Node {
	return e.parent
}

func (e *edgy) Child() *Node {
	return e.child
}

func (e *edgy) Value() interface{} {
	return e.value
}
