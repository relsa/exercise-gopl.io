package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)

	if n == 0 {
		return s
	}

	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + comma(s[1:])
	}

	if i := strings.IndexByte(s, '.'); i > 0 {
		return comma(s[:i]) + s[i:]
	}

	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
