// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

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
