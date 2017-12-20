package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestCountDiffBit(t *testing.T) {
	tests := []struct {
		ha   [32]byte
		hb   [32]byte
		want int
	}{
		{sha256.Sum256([]byte("hoge")), sha256.Sum256([]byte("hoge")), 0},
		{sha256.Sum256([]byte("fuga")), sha256.Sum256([]byte("fuga")), 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("countDiffBit(%x, %x)", test.ha, test.hb)
		got := countDiffBit(test.ha, test.hb)
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}
