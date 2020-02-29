package sock_merchant

import (
	"testing"
)

func TestSockMerchant(t *testing.T) {
	type test struct {
		n   int32
		ar  []int32
		out int32
	}

	testcases := []test{
		{n: 9, ar: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, out: 3},
	}

	t.Parallel()
	for _, c := range testcases {
		out := sockMerchant(c.n, c.ar)
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}
