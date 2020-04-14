package heap

type Heap struct {
	Items []int
}

func (h *Heap) LeftIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap) RightIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap) ParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap) HasLeft(index int) bool {
	return h.LeftIndex(index) < len(h.Items)
}

func (h *Heap) HasRight(index int) bool {
	return h.RightIndex(index) < len(h.Items)
}

func (h *Heap) HasParent(index int) bool {
	return h.ParentIndex(index) >= 0
}

func (h *Heap) LeftValue(index int) int {
	return h.Items[h.LeftIndex(index)]
}

func (h *Heap) RightValue(index int) int {
	return h.Items[h.RightIndex(index)]
}

func (h *Heap) ParentValue(index int) int {
	return h.Items[h.ParentIndex(index)]
}

func (h *Heap) Swap(indexOne, indexTwo int) {
	h.Items[indexOne], h.Items[indexTwo] = h.Items[indexTwo], h.Items[indexOne]
}
