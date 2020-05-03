// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package graph

import (
	"errors"
	"fmt"
	"sync"
)

type Heap struct {
	internal []interface{}

	// kv mapping of a node value to the node pointer
	nodes map[interface{}]*Node

	// kv mapping of a node value to it's index in the internal array
	values map[interface{}]int

	size int
	once sync.Once
}

func (h *Heap) Print() {
	for i := 0; i < h.size; i++ {
		v := h.nodes[h.internal[i]]
		fmt.Printf("index: %v | value: %v | cost: %v\n", i, v.Value, v.Cost)
	}
}

func (h *Heap) init() {
	h.once.Do(func() {
		h.internal = make([]interface{}, 0)
		h.nodes = make(map[interface{}]*Node)
		h.values = make(map[interface{}]int)
	})
}

func (h *Heap) Swap(i, j int) {
	h.init()

	iv := h.internal[i]
	jv := h.internal[j]

	// Swap the internal values for the dictionary
	h.internal[i] = jv
	h.values[jv] = i

	h.internal[j] = iv
	h.values[iv] = j
}

func (h *Heap) Up(i int) {
	h.init()

	if h.size > 0 {

		j := 0
		if i > 1 {
			j = i / 2 // Parent of i
		}

		if h.nodes[h.internal[i]].Cost < h.nodes[h.internal[j]].Cost {
			h.Swap(i, j)
			h.Up(j)
		}
	}
}

func (h *Heap) Down(i int) {
	h.init()

	if h.size > 0 {

		var j int
		if 2*i <= h.size {
			if 2*i < h.size {
				left := 2 * i
				right := (2 * i) + 1

				// set j as left, unless right is lower cost
				j = left

				if h.nodes[h.internal[right]] != nil && h.nodes[h.internal[right]].Cost < h.nodes[h.internal[left]].Cost {
					j = right
				}
			} else {
				j = 2 * i
			}

			if h.nodes[h.internal[j]] != nil {

				if h.nodes[h.internal[i]] == nil || h.nodes[h.internal[j]].Cost < h.nodes[h.internal[i]].Cost {
					h.Swap(i, j)
					h.Down(j)
				}
			}
		}
	}
}

func (h *Heap) Insert(n *Node) (err error) {
	h.init()

	// Set the dictionary value
	h.nodes[n.Value] = n

	// Add the value to the internal array
	h.internal = append(h.internal, n.Value)

	// Increment the size
	h.size++

	// add the map entry for the new element's index
	h.values[n.Value] = h.size - 1

	// Heap Up
	h.Up(h.size - 1)

	return err
}

func (h *Heap) Min() (index int, err error) {
	h.init()

	if h.size > 0 {

		// root node
		index = 0
	} else {
		// empty heap
		err = errors.New("empty heap")
	}

	return index, err
}

func (h *Heap) ExtractMin() (min *Node, err error) {
	h.init()

	var mini int
	if mini, err = h.Min(); err == nil {

		min = h.nodes[h.internal[mini]]
		h.Delete(mini)
	}

	return min, err
}

func (h *Heap) Delete(i int) (err error) {
	h.init()

	// Clean up maps
	v := h.internal[i]
	delete(h.nodes, v)
	delete(h.values, v)

	// Delete an entry from the internal slice while
	// maintaining it's ordering
	//copy(h.internal[i:], h.internal[i+1:])

	for index := i + 1; index < h.size; index++ {

		// Shift cells up the array
		v := h.internal[index]
		h.internal[index-1] = v
		h.values[v] = index - 1
		h.internal[index] = nil
	}

	// decrease the size now that the heap has been re-adjusted
	h.size--

	// h.internal[h.size-1] = ""
	// h.internal = h.internal[:len(h.internal)-1]

	// Reorder the heap from the root
	h.Down(0)

	return err
}

func (h *Heap) Find(value interface{}) (index int, err error) {

	return index, err
}

func (h *Heap) DeleteElem(value interface{}) (err error) {
	h.init()

	return err
}

func (h *Heap) ChangeCost(value interface{}, parent *Node, cost float64) {
	h.init()

	if h.nodes[value] != nil && cost < h.nodes[value].Cost {
		h.nodes[value].Parent = parent

		// Update the cost of the node
		h.nodes[value].Cost = cost

		// Move the node up the heap
		h.Up(h.values[value])
	}
}

func (h *Heap) Size() int {
	return h.size
}
