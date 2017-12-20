package main

import (
	"unicode/utf8"
)

func reverse(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	_, size := utf8.DecodeRune(b)

	rev(b[:size])
	rev(b[size:])
	rev(b)

	return reverse(b[:len(b)-size])

}

func rev(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
