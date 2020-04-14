package min_heap

import (
	"fmt"
	"sort"
	"testing"
)

func TestMinHeap(t *testing.T) {
	type test struct {
		initial []int
		toAdd   []int
	}

	testcases := []test{
		{[]int{}, []int{4, 9, 10, 0, -4, 7}},
		{[]int{}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{300, 5, 77, -8, 0, 50}},
		{[]int{}, []int{-1000, 1000}},
		{[]int{}, []int{1000, -1000}},
		{[]int{4, 9, 10, 0, -4, 7}, []int{}},
		{[]int{0, 7, 10}, []int{1, 2, 3, 4, 5}},
		{[]int{100}, []int{300, 5, 77, -8, 0, 50}},
		{[]int{-2000, 0, 800}, []int{-1000, 1000}},
		{[]int{5000, 10000}, []int{1000, -1000}},
	}

	for i, test := range testcases {
		h := New(test.initial)

		for _, n := range test.toAdd {
			h.insert(n)
		}

		checkMinHeap(i, h, t)
	}
}

func checkMinHeap(n int, h *minHeap, t *testing.T) {
	result := make([]int, 0)

	correctlySorted := make([]int, len(h.Items))
	copy(correctlySorted, h.Items)

	length := len(h.Items)
	for i := 0; i < length; i++ {
		result = append(result, h.ExtractMin())
	}

	sort.Slice(correctlySorted, func(i, j int) bool {
		return correctlySorted[i] < correctlySorted[j]
	})

	for k := range correctlySorted {
		if result[k] != correctlySorted[k] {
			t.Errorf("%v does not equal %v", result, correctlySorted)
			return
		}
	}

	fmt.Printf("[%v] correctly sorted: %v\n", n, result)
}
