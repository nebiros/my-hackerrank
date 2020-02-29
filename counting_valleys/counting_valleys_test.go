package counting_valleys

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	type test struct {
		n   int32
		s  string
		out int32
	}

	testcases := []test{
		{n: 8, s: "UDDDUDUU", out: 1},
		{n: 12, s: "DDUUDDUDUUUD", out: 2},
	}

	t.Parallel()
	for _, c := range testcases {
		out := countingValleys(c.n, c.s)
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}