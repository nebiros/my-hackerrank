package jumping_on_the_clouds

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	type test struct {
		c   []int32
		out int32
	}

	testcases := []test{
		{c: []int32{0, 0, 1, 0, 0, 1, 0}, out: 4},
		{c: []int32{0, 0, 0, 1, 0, 0}, out: 3},
	}

	t.Parallel()
	for _, c := range testcases {
		out := jumpingOnClouds(c.c)
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}
