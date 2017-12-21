package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type unicodeCtg string

const (
	Letter unicodeCtg = "Letter"
	Number            = "Number"
	Space             = "Space"
)

var ctgFunctions = map[unicodeCtg]func(rune) bool{
	Letter: unicode.IsLetter,
	Number: unicode.IsNumber,
	Space:  unicode.IsSpace,
}

func main() {
	counts, invalid := countUnicodeCtg(os.Stdin)
	fmt.Printf("ctg\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func countUnicodeCtg(r io.Reader) (map[unicodeCtg]int, int) {
	counts := make(map[unicodeCtg]int)
	invalid := 0

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		for c, f := range ctgFunctions {
			if f(r) {
				counts[c]++
			}
		}
	}
	return counts, invalid
}
