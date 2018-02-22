package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(NewReader(os.Args[1]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

type reader struct {
	s string
	n int
}

func (r *reader) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(r.s)[r.n:])
	r.n += n
	if r.n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

func NewReader(s string) io.Reader {
	return &reader{s: s}
}
