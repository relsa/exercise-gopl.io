package main

import (
	"fmt"
	"io"
	"strings"
)

func EchoA(out io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func EchoB(out io.Writer, args []string) {
	fmt.Fprintln(out, strings.Join(args[1:], " "))
}
