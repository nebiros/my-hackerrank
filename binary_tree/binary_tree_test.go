package binary_tree

import (
	"testing"

	"github.com/pkg/errors"
)

func TestTreeFind(t *testing.T) {
	tr := &tree{root: nil}

	seed := map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}

	for k, v := range seed {
		err := tr.insert(k, v)
		if err != nil {
			t.Fatalf("cannot insert value '%s': %s", k, err.Error())
		}
	}

	type test struct {
		value string
		out   struct {
			data  string
			found bool
		}
	}

	testcases := []test{
		{value: "d", out: struct {
			data  string
			found bool
		}{data: "delta", found: true}},
		{value: "b", out: struct {
			data  string
			found bool
		}{data: "bravo", found: true}},
		{value: "c", out: struct {
			data  string
			found bool
		}{data: "charlie", found: true}},
		{value: "e", out: struct {
			data  string
			found bool
		}{data: "echo", found: true}},
		{value: "a", out: struct {
			data  string
			found bool
		}{data: "alpha", found: true}},
		{value: "x", out: struct {
			data  string
			found bool
		}{data: "", found: false}},
		{value: "y", out: struct {
			data  string
			found bool
		}{data: "", found: false}},
		{value: "z", out: struct {
			data  string
			found bool
		}{data: "", found: false}},
	}

	t.Parallel()

	for _, c := range testcases {
		data, found := tr.find(c.value)
		if found != c.out.found {
			t.Errorf("'%s' not found", c.value)
			continue
		}

		if data != c.out.data {
			t.Errorf("'%s' data different than '%s'", c.value, c.out.data)
			continue
		}
	}
}

func TestTreeDelete(t *testing.T) {
	type test struct {
		seed  map[string]string
		value string
		out   error
	}

	testcases := []test{
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "d", out: nil},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "b", out: nil},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "c", out: nil},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "e", out: nil},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "a", out: nil},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "x", out: new(errTreeNodeNil)},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "y", out: new(errTreeNodeNil)},
		{seed: map[string]string{"d": "delta", "b": "bravo", "c": "charlie", "e": "echo", "a": "alpha"}, value: "z", out: new(errTreeNodeNil)},
	}

	t.Parallel()

	for _, c := range testcases {
		tr := &tree{root: nil}

		for k, v := range c.seed {
			err := tr.insert(k, v)
			if err != nil {
				t.Fatalf("cannot insert value '%s': %s", k, err.Error())
			}
		}

		err := tr.delete(c.value)
		if err != nil {
			if errors.Cause(err) != c.out {
				t.Errorf("cannot delete '%s': %s", c.value, err.Error())
			}
		}
	}
}
