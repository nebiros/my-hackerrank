package longest_len_bi_valued_slice

import "testing"

func Test_longestLen(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"whole array is bi-valued": {
			input: []int{4, 2, 2, 4, 2},
			want:  5,
		},
		"longest bi-valued slice 1": {
			input: []int{1, 2, 3, 2},
			want:  3,
		},
		"longest bi-valued slice 2": {
			input: []int{0, 5, 4, 4, 5, 1},
			want:  4,
		},
		"longest bi-valued slice 3": {
			input: []int{4, 4},
			want:  2,
		},
	}

	t.Parallel()

	for name, tc := range tests {
		tc := tc

		t.Run(name, func(t *testing.T) {
			got := longestLen(tc.input)
			if got != tc.want {
				t.Errorf("longestLen() = %v, want %v", got, tc.want)
			}
		})
	}
}
