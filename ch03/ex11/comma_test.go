package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"+1", "+1"},
		{"-1", "-1"},
		{"1234.5678", "1,234.5678"},
		{"-1234.5678", "-1,234.5678"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("comma(%q)", test.s)
		got := comma(test.s)
		want := test.want

		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}
