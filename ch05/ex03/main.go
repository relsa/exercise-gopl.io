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
	printTextNode(doc)
}

func printTextNode(n *html.Node) {
	if n == nil {
		return
	}

	switch n.Type {
	case html.TextNode:
		fmt.Println(n.Data)
	case html.ElementNode:
		if n.Data == "script" || n.Data == "style" {
			break
		}
		printTextNode(n.FirstChild)
	default:
		printTextNode(n.FirstChild)
	}

	printTextNode(n.NextSibling)
}
