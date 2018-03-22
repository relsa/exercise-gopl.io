package strings_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	ts := []struct {
		s, sep string
		want   int
	}{
		{"", ":", 1},
		{"a:b:c", ":", 3},
		{"a b c d", " ", 4},
	}
	for _, tc := range ts {
		words := strings.Split(tc.s, tc.sep)
		if got, want := len(words), tc.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", tc.s, tc.sep, got, want)
		}
	}
}
