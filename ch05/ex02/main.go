package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	countTag(counts, doc)

	for k, v := range counts {
		fmt.Printf("%s:\t%d\n", k, v)
	}
}

func countTag(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countTag(counts, c)
	}
}
