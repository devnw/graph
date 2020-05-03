// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package graph

type dedgy struct {
	parent *Node
	child  *Node
	value  interface{}
}

func (e *dedgy) Parent() *Node {
	return e.parent
}

func (e *dedgy) Child() *Node {
	return e.child
}

func (e *dedgy) Value() interface{} {
	return e.value
}

func (e *dedgy) Directed() bool {
	return true
}
