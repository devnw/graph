package graph

type edgy struct {
	parent Node
	child  Node
	value  interface{}
}

func (e *edgy) Parent() Node {
	return e.parent
}

func (e *edgy) Child() Node {
	return e.child
}

func (e *edgy) Value() interface{} {
	return e.value
}

type directedEdgy struct {
	parent Node
	child  Node
	value  interface{}
}

func (e *directedEdgy) Parent() Node {
	return e.parent
}

func (e *directedEdgy) Child() Node {
	return e.child
}

func (e *directedEdgy) Value() interface{} {
	return e.value
}

func (e *directedEdgy) Directed() bool {
	return true
}

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

type directedWedgy struct {
	parent Node
	child  Node
	value  interface{}
	weight int
}

func (e *directedWedgy) Parent() Node {
	return e.parent
}

func (e *directedWedgy) Child() Node {
	return e.child
}

func (e *directedWedgy) Value() interface{} {
	return e.value
}

func (e *directedWedgy) Weight() int {
	return e.weight
}

func (e *directedWedgy) Directed() bool {
	return true
}
