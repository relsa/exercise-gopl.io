package params

import (
	"testing"
)

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func TestPack(t *testing.T) {
	d := data{
		Labels:     []string{"hoge", "fuga"},
		MaxResults: 10,
		Exact:      true,
	}
	want := "l=hoge&l=fuga&max=10&x=true"

	got := Pack(&d)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
