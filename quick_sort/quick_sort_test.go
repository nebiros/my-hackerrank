package quick_sort

import (
	"reflect"
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
		{arr: []int{}, out: []int{}},
		{arr: []int{42}, out: []int{42}},
		{arr: []int{42, 23}, out: []int{23, 42}},
		{arr: []int{23, 42, 32, 64, 12, 4}, out: []int{4, 12, 23, 32, 42, 64}},
	}

	t.Parallel()

	for _, c := range testcases {
		quickSort(c.arr, 0, len(c.arr)-1)

		if !sort.IntsAreSorted(c.arr) {
			t.Errorf("%v not sorted", c.arr)
			continue
		}

		actual := c.arr
		expected := c.out

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v != %v", actual, expected)
			continue
		}
	}
}
