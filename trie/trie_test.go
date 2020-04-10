package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tr := &trie{root: &trieNode{}}

	seed := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(seed); i++ {
		tr.insert(seed[i])
	}

	type test struct {
		w string
		out bool
	}

	testcases := []test{
		{w: "sam", out: true},
		{w: "john", out: true},
		{w: "tim", out: true},
		{w: "jose", out: true},
		{w: "rose", out: true},
		{w: "cat", out: true},
		{w: "dog", out: true},
		{w: "dogg", out: true},
		{w: "roses", out: true},
		{w: "rosess", out: false},
		{w: "ans", out: false},
		{w: "san", out: false},
	}

	t.Parallel()

	for _, c := range testcases {
		out := tr.find(c.w)
		if out != c.out {
			t.Errorf("'%s' not found", c.w)
		}
	}
}
