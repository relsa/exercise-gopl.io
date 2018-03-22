package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	ts := []struct {
		input   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			input:   nil,
			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("a"),
			counts:  map[rune]int{'a': 1},
			utflen:  []int{0, 1, 0, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("à"),
			counts:  map[rune]int{'à': 1},
			utflen:  []int{0, 0, 1, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("あ"),
			counts:  map[rune]int{'あ': 1},
			utflen:  []int{0, 0, 0, 1, 0},
			invalid: 0,
		},
		{
			input:   []byte("🍺"),
			counts:  map[rune]int{'🍺': 1},
			utflen:  []int{0, 0, 0, 0, 1},
			invalid: 0,
		},
		{
			input:   []byte("aあa"),
			counts:  map[rune]int{'a': 2, 'あ': 1},
			utflen:  []int{0, 2, 0, 1, 0},
			invalid: 0,
		},
		{
			input:   []byte{0xfe},
			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 1,
		},
	}
	for _, tc := range ts {
		counts, utflen, invalid, err := CharCount(bytes.NewReader(tc.input))
		if err != nil {
			t.Error(err)
			continue
		}
		if got, want := counts, tc.counts; !reflect.DeepEqual(got, want) {
			t.Errorf("counts want: %v, but got: %v", want, got)
		}
		if got, want := utflen, tc.utflen; !reflect.DeepEqual(got, want) {
			t.Errorf("utflen want: %v, but got: %v", want, got)
		}
		if got, want := invalid, tc.invalid; got != want {
			t.Errorf("invalid want: %v, but got: %v", want, got)
		}
	}
}
