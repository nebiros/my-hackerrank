package quick_sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type test struct {
		arr []int
		out []int
	}

	testcases := []test{
		{arr: []int{15, 3, 12, 6, -9, 9, 0}, out: []int{-9, 0, 3, 6, 9, 12, 15}},
	}

	t.Parallel()

	for _, c := range testcases {
		quickSort(c.arr, 0, len(c.arr)-1)

		if !sort.IntsAreSorted(c.arr) {
			t.Errorf("%v not sorted", c.arr)
		}
	}
}
