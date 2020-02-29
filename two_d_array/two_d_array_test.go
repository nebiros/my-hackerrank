package two_d_array

import (
	"testing"
)

func TestHourglassSum(t *testing.T) {
	type test struct {
		arr [][]int32
		out int32
	}

	testcases := []test{
		{
			arr: [][]int32{{1, 1, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0}, {0, 0, 2, 4, 4, 0}, {0, 0, 0, 2, 0, 0}, {0, 0, 1, 2, 4, 0}},
			out: 19,
		},
		{
			arr: [][]int32{{1, 1, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0}, {0, 9, 2, -4, -4, 0}, {0, 0, 0, -2, 0, 0}, {0, 0, -1, -2, -4, 0}},
			out: 13,
		},
		{
			arr: [][]int32{{-9, -9, -9, 1, 1, 1}, {0, -9, 0, 4, 3, 2}, {-9, -9, -9, 1, 2, 3}, {0, 0, 8, 6, 6, 0}, {0, 0, 0, -2, 0, 0}, {0, 0, 1, 2, 4, 0}},
			out: 28,
		},
	}

	t.Parallel()
	for _, c := range testcases {
		out := hourglassSum(c.arr)
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}
