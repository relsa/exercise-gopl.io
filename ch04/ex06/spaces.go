package main

import (
	"unicode"
	"unicode/utf8"
)

func rmAdjacentSpaces(b []byte) []byte {

	if len(b) == 0 {
		return b
	}

	var r rune
	var size int
	var ptr int
	var skip bool

	for i := 0; i < len(b); i += size {
		r, size = utf8.DecodeRune(b[i:])

		if unicode.IsSpace(r) {
			if !skip {
				buf := make([]byte, 4)
				s := utf8.EncodeRune(buf, ' ')

				write(buf, b, 0, ptr, size)
				ptr += s
				skip = true
			}
			continue
		}

		write(b, b, i, ptr, size)
		ptr += size
		skip = false
	}

	return b[:ptr]
}

func write(src, dest []byte, sptr, dptr, size int) {
	for i := 0; i < size; i++ {
		dest[dptr+i] = src[sptr+i]
	}
}
