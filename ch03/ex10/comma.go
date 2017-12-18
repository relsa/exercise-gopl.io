package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", Comma(os.Args[i]))
	}
}

func Comma(s string) string {
	b := []byte(s)

	var buf bytes.Buffer

	n := len(b)
	m := n % 3

	for i := 0; i < n; i++ {
		if (m == 0 && i > 0 && i%3 == 0) ||
			(m > 0 && (i-m)%3 == 0) {
			buf.WriteByte(',')
		}
		buf.WriteByte(b[i])
	}

	return buf.String()
}
