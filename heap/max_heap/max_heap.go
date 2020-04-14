package max_heap

import (
	"log"

	"github.com/nebiros/my-hackerrank/heap"
)

type maxHeap struct {
	*heap.Heap
}

func New(items []int) *maxHeap {
	h := &maxHeap{
		Heap: &heap.Heap{
			Items: items,
		},
	}

	if len(h.Items) > 0 {
		h.buildMaxHeap()
	}

	return h
}

func (h *maxHeap) buildMaxHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

func (h *maxHeap) insert(item int) *maxHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.maxHeapifyUp(lastElementIndex)

	return h
}

func (h *maxHeap) maxHeapifyUp(index int) {
	for h.HasParent(index) && h.ParentValue(index) < h.Items[index] {
		h.Swap(h.ParentIndex(index), index)
		index = h.ParentIndex(index)
	}
}

func (h *maxHeap) extractMax() int {
	if len(h.Items) == 0 {
		log.Fatal("no items in the heap")
	}

	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	// shrinking slice
	h.Items = h.Items[:len(h.Items)-1]

	h.maxHeapifyDown(0)

	return minItem
}

func (h *maxHeap) maxHeapifyDown(index int) {
	// iterate until we have a child node smaller than the index value
	for (h.HasLeft(index) && h.Items[index] < h.LeftValue(index)) ||
		(h.HasRight(index) && h.Items[index] < h.RightValue(index)) {
		// if both children are smaller
		if (h.HasLeft(index) && h.Items[index] < h.LeftValue(index)) &&
			(h.HasRight(index) && h.Items[index] < h.RightValue(index)) {
			// compare the children and swap with the largest
			if h.LeftValue(index) > h.RightValue(index) {
				h.Swap(index, h.LeftIndex(index))
				index = h.LeftIndex(index)
			} else {
				h.Swap(index, h.RightIndex(index))
				index = h.RightIndex(index)
			}
		} else if h.HasLeft(index) && h.Items[index] < h.LeftValue(index) { // else if only left child is smaller than swap with it
			h.Swap(index, h.LeftIndex(index))
			index = h.LeftIndex(index)
		} else { // else it must be the right child that is smaller, so we swap with it
			h.Swap(index, h.RightIndex(index))
			index = h.RightIndex(index)
		}
	}
}
