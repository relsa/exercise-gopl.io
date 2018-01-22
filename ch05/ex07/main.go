package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := outline(url); err != nil {
			log.Print(err)
		}
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		var attrStr string
		for _, a := range n.Attr {
			attrStr += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrStr)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrStr)
			depth++
		}

	case html.TextNode:
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			fmt.Printf("%*s%s\n", depth*2, "", s)
		}
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
