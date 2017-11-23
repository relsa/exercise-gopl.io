package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	Echo(out, []string{"./echo", "go", "programming", "language"})

	got := out.String()
	want := "./echo go programming language\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
