package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		arr  [8]int
		want [8]int
	}{
		{[8]int{0, 1, 2, 3, 4, 5, 6, 7}, [8]int{7, 6, 5, 4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		reverse(&test.arr)
		for i := 0; i < 8; i++ {
			if test.arr[i] != test.want[i] {
				t.Errorf("arr[%d] = %d, want %d", i, test.arr[i], test.want[i])
			}
		}
	}
}
