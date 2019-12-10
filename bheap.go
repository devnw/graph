package graph

import (
	"errors"
	"sync"
)

type heap struct {
	internal []interface{}
	dict     map[interface{}]*Node
	size     int
	once     sync.Once
}

func (h *heap) init() {
	h.once.Do(func() {
		h.internal = make([]interface{}, 0)
		h.dict = make(map[interface{}]*Node)
	})
}

func (h *heap) Swap(i, j int) {
	h.init()

	iv := h.internal[i]
	jv := h.internal[j]

	// Swap the internal values for the dictionary
	h.internal[i] = jv
	h.internal[j] = iv
}

func (h *heap) Up(i int) {
	h.init()

	if i > 1 {
		j := i / 2 // Parent of i
		if h.dict[h.internal[i]].Cost < h.dict[h.internal[j]].Cost {
			h.Swap(i, j)
			h.Up(j)
		}
	}
}

func (h *heap) Down(i int) {
	h.init()

	var j int
	if 2*i <= h.size {
		if 2*i < h.size {
			left := 2 * i
			right := (2 * i) + 1

			// set j as left, unless right is lower cost
			j = left
			if h.dict[h.internal[right]].Cost < h.dict[h.internal[left]].Cost {
				j = right
			}
		} else {
			j = 2 * i
		}

		if h.dict[h.internal[j]].Cost < h.dict[h.internal[i]].Cost {
			h.Swap(i, j)
			h.Down(j)
		}
	}
}

func (h *heap) Insert(n *Node) (err error) {
	h.init()

	// Set the dictionary value
	h.dict[n.Value] = n

	// Add the value to the internal array
	h.internal = append(h.internal, n.Value)

	// Increment the size
	h.size++

	// Heap Up
	h.Up(h.size - 1)

	return err
}

func (h *heap) Min() (index int, err error) {
	h.init()

	if h.size > 0 {
		if h.size == 1 {
			// root node
			index = 0
		} else {
			index = h.size - 1
		}
	} else {
		// empty heap
		err = errors.New("empty heap")
	}

	return index, err
}

func (h *heap) ExtractMin() (min *Node, err error) {
	h.init()

	var mini int
	if mini, err = h.Min(); err == nil {

		min = h.dict[h.internal[mini]]
		h.Delete(mini)
	}

	return min, err
}

func (h *heap) Delete(i int) (err error) {
	h.init()

	// Delete an entry from the internal slice while
	// maintaining it's ordering
	copy(h.internal[i:], h.internal[i+1:])
	h.internal[h.size-1] = ""
	h.internal = h.internal[:len(h.internal)-1]

	// Reorder the heap from the root
	h.Down(0)

	return err
}

func (h *heap) Find(value interface{}) (index int, err error) {

	return index, err
}

func (h *heap) DeleteElem(value interface{}) (err error) {
	h.init()

	return err
}

func (h *heap) SetKey(old, new interface{}) (err error) {
	h.init()

	return err
}
