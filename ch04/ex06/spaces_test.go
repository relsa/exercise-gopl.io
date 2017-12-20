package main

import (
	"fmt"
	"testing"
)

func TestRmAdjacentDups(t *testing.T) {
	tests := []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("ab"), []byte("ab")},
		{[]byte("a b"), []byte("a b")},
		{[]byte("a　b"), []byte("a b")},
		{[]byte("a　　　b"), []byte("a b")},
		{[]byte("阿　　　吽"), []byte("阿 吽")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s", string(test.b))
		got := rmAdjacentSpaces(test.b)
		if string(got) != string(test.want) {
			t.Errorf("case %q: got %q, want %q", descr, string(got), string(test.want))
		}
	}
}
