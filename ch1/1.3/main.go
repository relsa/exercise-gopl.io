// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	measure(echoA)
	measure(echoB)
}

func measure(f func()) {
	start := time.Now()

	f()

	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

func echoA() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoB() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//!-
