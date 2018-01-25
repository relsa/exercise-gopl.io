package main

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	ts := []struct {
		s string
		n int64
	}{
		{"", 3},
		{"hogehoge", 4},
		{"hoge", 8},
		{"日本語", 6},
	}

	for _, tc := range ts {

		r1 := strings.NewReader(tc.s)
		r2 := strings.NewReader(tc.s)
		b1 := make([]byte, tc.n)
		b2 := make([]byte, tc.n)

		lr1 := LimitReader(r1, tc.n)
		lr2 := io.LimitReader(r2, tc.n)

		n1, err1 := lr1.Read(b1)
		n2, err2 := lr2.Read(b2)
		if err1 != err2 {
			t.Errorf("err1=%v, err2=%v", err1, err2)
		}
		if n1 != n2 {
			t.Errorf("n1=%d, n2=%d", n1, n2)
		}
	}
}
