package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		sa   string
		sb   string
		want bool
	}{
		{"", "", true},
		{"hoge", "egoh", true},
		{"hoge", "hogehoge", false},
		{"digdag", "gdgdai", true},
		{"digdag", "gdgd", false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("anagram(%q, %q)", test.sa, test.sb)
		got := anagram(test.sa, test.sb)
		want := test.want

		if got != test.want {
			t.Errorf("%s = %t, want %t", descr, got, want)
		}
	}
}
