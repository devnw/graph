package graph

type dwedgy struct {
	parent Node
	child  Node
	value  interface{}
	weight int
}

func (e *dwedgy) Parent() Node {
	return e.parent
}

func (e *dwedgy) Child() Node {
	return e.child
}

func (e *dwedgy) Value() interface{} {
	return e.value
}

func (e *dwedgy) Weight() int {
	return e.weight
}

func (e *dwedgy) Directed() bool {
	return true
}
