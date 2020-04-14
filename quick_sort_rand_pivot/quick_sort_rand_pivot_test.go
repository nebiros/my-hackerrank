package quick_sort_rand_pivot

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
		actual := quickSort(c.arr)
		expected := c.out

		if !sort.IntsAreSorted(actual) {
			t.Errorf("%v not sorted", actual)
			continue
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v != %v", actual, expected)
			continue
		}
	}
}
