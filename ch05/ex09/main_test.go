package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"", strings.ToUpper, ""},
		{"$", strings.ToUpper, ""},
		{"$hoge", strings.ToUpper, "HOGE"},
		{"$hoge fuga", strings.ToUpper, "HOGE fuga"},
		{"$hoge $fuga", strings.ToUpper, "HOGE FUGA"},
	}

	for _, test := range tests {
		got := expand(test.s, test.f)
		if got != test.want {
			t.Fatalf("case %q: got %q, want %q", test.s, got, test.want)
		}
	}
}
