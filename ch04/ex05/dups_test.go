package main

import (
	"fmt"
	"testing"
)

func TestRmAdjacentDups(t *testing.T) {
	tests := []struct {
		strings []string
		want    []string
	}{
		{nil, nil},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "a", "a"}, []string{"a"}},
		{[]string{"a", "a", "a", "b", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v", test.strings)
		got := rmAdjacentDups(test.strings)
		if len(got) != len(test.want) {
			t.Errorf("len(%v) = %d, want %d", got, len(got), len(test.want))
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("case %s: got[%d] = %q, want %q", descr, i, got[i], test.want[i])
			}
		}
	}
}
