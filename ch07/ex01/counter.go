package main

import (
	"bufio"
)

type TermCounter int

func (c *TermCounter) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanWords(p, true)

		if err != nil {
			return 0, err
		}
		if token != nil {
			*c++
		}

		p = p[advance:]
	}

	return n, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanLines(p, true)

		if err != nil {
			return 0, err
		}
		if token != nil {
			*c++
		}

		p = p[advance:]
	}

	return n, nil
}
