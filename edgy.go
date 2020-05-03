// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

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
