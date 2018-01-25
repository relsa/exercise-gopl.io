package main

import (
	"testing"
)

func TestTermCounter(t *testing.T) {
	ts := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"hoge", 1},
		{"hoge  fuga\npiyo\n", 3},
	}

	for _, tc := range ts {
		var c TermCounter
		b := []byte(tc.s)
		n, err := c.Write(b)

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if n != len(b) {
			t.Errorf("written byte size is %d, want %d", n, len(b))
		}

		if c != TermCounter(tc.want) {
			t.Errorf("count is %d, want %d", c, tc.want)
		}
	}
}
func TestLineCounter(t *testing.T) {
	ts := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"hoge", 1},
		{"hoge\nfuga\npiyo\n", 3},
	}

	for _, tc := range ts {
		var c LineCounter
		b := []byte(tc.s)
		n, err := c.Write(b)

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if n != len(b) {
			t.Errorf("written byte size is %d, want %d", n, len(b))
		}

		if c != LineCounter(tc.want) {
			t.Errorf("count is %d, want %d", c, tc.want)
		}
	}
}
