package ctci_array_left_rotation

import (
	"reflect"
	"testing"
)

func TestRotLeft(t *testing.T) {
	type test struct {
		a   []int32
		d   int32
		out []int32
	}

	testcases := []test{
		{
			a:   []int32{1, 2, 3, 4, 5},
			d:   4,
			out: []int32{5, 1, 2, 3, 4},
		},
		{
			a:   []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
			d:   10,
			out: []int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
		},
		{
			a:   []int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97},
			d:   13,
			out: []int32{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60},
		},
	}

	t.Parallel()
	for _, c := range testcases {
		out := rotLeft(c.a, c.d)
		if !reflect.DeepEqual(out, c.out) {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}
