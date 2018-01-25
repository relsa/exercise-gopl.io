package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	ts := []struct {
		s string
	}{
		{""},
		{"hoge"},
		{"hoge fuga piyo"},
		{"ほげ"},
	}

	for _, tc := range ts {
		var b bytes.Buffer

		cw, n := CountingWriter(&b)
		if *n != 0 {
			t.Errorf("got count=%d, want count=%d", *n, 0)
		}
		io.WriteString(cw, tc.s)

		if *n != int64(len(tc.s)) {
			t.Errorf("got count=%d, want count=%d", *n, len(tc.s))
		}

		if got := b.String(); got != tc.s {
			t.Errorf("got out=%d, want out=%d", got, tc.s)
		}
	}
}
