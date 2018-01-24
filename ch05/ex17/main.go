package main

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var hit []*html.Node

	forEachNode(doc, func(n *html.Node) {

		if n.Type == html.ElementNode {
			for _, nm := range name {
				if n.Data == nm {
					hit = append(hit, n)
					break
				}
			}
		}

	}, nil)

	return hit
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
