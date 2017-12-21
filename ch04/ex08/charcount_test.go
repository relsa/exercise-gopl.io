package main

import (
	"bytes"
	"testing"
)

func TestCountUnicodeCtg(t *testing.T) {
	tests := []struct {
		in   string
		want map[unicodeCtg]int
	}{
		{
			"hoge fuga 1234",
			map[unicodeCtg]int{
				Letter: 8, Number: 4, Space: 2,
			},
		},
	}

	for _, test := range tests {
		in := bytes.NewBufferString(test.in)
		got, _ := countUnicodeCtg(in)

		for k, v := range got {
			if w, ok := test.want[k]; !ok || v != w {
				t.Errorf("case %q: got %v, want %v", test.in, got, test.want)
			}
		}

	}
}
