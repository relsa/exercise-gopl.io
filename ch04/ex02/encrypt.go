package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var alg = flag.String("x", "sha256", "sha256/384/512")

func main() {
	flag.Parse()

	switch *alg {
	case "sha256", "sha384", "sha512":
		// pass
	default:
		fmt.Fprintf(os.Stderr, "invalid algorithm: %s\n", *alg)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d := scanner.Bytes()
		switch *alg {
		case "sha256":
			res := sha256.Sum256(d)
			fmt.Printf("%x\n", res)
		case "sha384":
			res := sha512.Sum384(d)
			fmt.Printf("%x\n", res)
		case "sha512":
			res := sha512.Sum512(d)
			fmt.Printf("%x\n", res)
		}
	}
}
