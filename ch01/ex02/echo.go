package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(out io.Writer, args []string) {
	for i, arg := range args[1:] {
		fmt.Fprintln(out, i, arg)
	}
}

func main() {
	Echo(os.Stdin, os.Args)
}
