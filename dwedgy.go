// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package graph

type dwedgy struct {
	parent *Node
	child  *Node
	value  interface{}
	weight float64
}

func (e *dwedgy) Parent() *Node {
	return e.parent
}

func (e *dwedgy) Child() *Node {
	return e.child
}

func (e *dwedgy) Value() interface{} {
	return e.value
}

func (e *dwedgy) Weight() float64 {
	return e.weight
}

func (e *dwedgy) Directed() bool {
	return true
}
