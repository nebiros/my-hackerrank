package min_heap

import (
	"log"

	"github.com/nebiros/my-hackerrank/heap"
)

type minHeap struct {
	*heap.Heap
}

func New(items []int) *minHeap {
	h := &minHeap{
		Heap: &heap.Heap{
			Items: items,
		},
	}

	if len(h.Items) > 0 {
		h.buildMinHeap()
	}

	return h
}

func (h *minHeap) buildMinHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.minHeapifyDown(i)
	}
}

func (h *minHeap) insert(item int) *minHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.minHeapifyUp(lastElementIndex)

	return h
}

func (h *minHeap) minHeapifyUp(index int) {
	for h.HasParent(index) && (h.ParentValue(index) > h.Items[index]) {
		h.Swap(h.ParentIndex(index), index)
		index = h.ParentIndex(index)
	}
}

func (h *minHeap) ExtractMin() int {
	if len(h.Items) == 0 {
		log.Fatal("no items in the heap")
	}

	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	// shrinking slice
	h.Items = h.Items[:len(h.Items)-1]

	h.minHeapifyDown(0)

	return minItem
}

func (h *minHeap) minHeapifyDown(index int) {
	for (h.HasLeft(index) && (h.Items[index] > h.LeftValue(index))) ||
		(h.HasRight(index) && (h.Items[index] > h.RightValue(index))) {
		if (h.HasLeft(index) && (h.Items[index] > h.LeftValue(index))) &&
			(h.HasRight(index) && (h.Items[index] > h.RightValue(index))) {
			if h.LeftValue(index) < h.RightValue(index) {
				h.Swap(index, h.LeftIndex(index))
				index = h.LeftIndex(index)
			} else {
				h.Swap(index, h.RightIndex(index))
				index = h.RightIndex(index)
			}
		} else if h.HasLeft(index) && (h.Items[index] > h.LeftValue(index)) {
			h.Swap(index, h.LeftIndex(index))
			index = h.LeftIndex(index)
		} else {
			h.Swap(index, h.RightIndex(index))
			index = h.RightIndex(index)
		}
	}
}
