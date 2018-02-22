package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	ts := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 2},
	}
	for _, tc := range ts {
		if got := PopCount(uint64(tc.n)); got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
