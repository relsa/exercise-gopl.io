package main

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		a    []string
		sep  string
		want string
	}{
		{nil, ",", ""},
		{[]string{"hoge"}, ",", "hoge"},
		{[]string{"hoge", "fuga"}, ",", "hoge,fuga"},
		{[]string{"hoge", "fuga"}, ";", "hoge;fuga"},
	}

	for _, test := range tests {
		got := Join(test.sep, test.a...)
		if got != test.want {
			fmt.Errorf("got %q, want %q", got, test.want)
		}
	}
}
