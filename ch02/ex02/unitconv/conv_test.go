package unitconv

import (
	"testing"
)

func TestCToK(t *testing.T) {
	x := AbsoluteZeroC
	got := CToK(x)
	want := Kelvin(0)
	if got != want {
		t.Errorf("CToK(%s) = %s, want %s", x, got, want)
	}
}

func TestKToC(t *testing.T) {
	x := Kelvin(0)
	got := KToC(x)
	want := AbsoluteZeroC
	if got != want {
		t.Errorf("KToC(%s) = %s, want %s", x, got, want)
	}
}
