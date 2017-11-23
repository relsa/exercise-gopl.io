package popcount

import (
	"gopl.io/ch2/popcount"
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{uint64(0), 0},
		{uint64(3), 2},
		{uint64(4), 1},
		{uint64(63), 6},
	}

	for _, test := range tests {
		if got := PopCount(test.x); got != test.want {
			t.Errorf("PopCount(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountOrigin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}
