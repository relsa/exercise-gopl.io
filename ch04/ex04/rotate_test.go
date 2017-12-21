package main

import (
	"fmt"
	"testing"
)

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	tests := []struct {
		s    []int
		n    int
		want []int
	}{
		{
			[]int{0, 1, 2, 3, 4},
			0,
			[]int{0, 1, 2, 3, 4},
		},
		{
			[]int{0, 1, 2, 3, 4},
			1,
			[]int{4, 0, 1, 2, 3},
		},
		{
			[]int{0, 1, 2, 3, 4},
			2,
			[]int{3, 4, 0, 1, 2},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v", test.s)
		rotate(test.s, test.n)
		if !equals(test.s, test.want) {
			t.Errorf("case %s: got %v, want %v", descr, test.want, test.s)
		}
	}
}
