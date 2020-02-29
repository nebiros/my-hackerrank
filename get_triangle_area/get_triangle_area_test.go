package get_triangle_area

import (
	"reflect"
	"testing"
)

func TestGetTriangleArea(t *testing.T) {
	type test struct {
		x   []int32
		y   []int32
		out int64
	}

	testcases := []test{
		test{x: []int32{0, 3, 0}, y: []int32{0, 5, 2}, out: 3},
	}

	t.Parallel()
	for _, c := range testcases {
		out := getTriangleArea(c.x, c.y)
		if !reflect.DeepEqual(out, c.out) {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}
