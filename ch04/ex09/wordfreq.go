package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		for k, v := range wordfreq(arg) {
			fmt.Printf("%q\t%d\n", k, v)
		}
	}
}

func wordfreq(fname string) map[string]int {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	defer f.Close()

	in := bufio.NewScanner(f)
	in.Split(bufio.ScanWords)

	tf := make(map[string]int)

	for in.Scan() {
		tf[in.Text()]++
	}

	return tf
}
