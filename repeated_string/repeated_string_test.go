package repeated_string

import "testing"

func TestRepeatedString(t *testing.T) {
	type test struct {
		s   string
		n   int64
		out int64
	}

	testcases := []test{
		{s: "aba", n: 10, out: 7},
		{s: "a", n: 1000000000000, out: 1000000000000},
	}

	t.Parallel()
	for _, c := range testcases {
		out := repeatedString(c.s, c.n)
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}
