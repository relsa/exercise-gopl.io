package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)
var maxDepth = flag.Int("depth", 0, "max depth from root url")

type node struct {
	depth int
	links []string
}

func crawl(depth int, url string) *node {
	if depth >= *maxDepth {
		return &node{depth + 1, nil}
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return &node{depth + 1, list}
}

func main() {
	flag.Parse()
	worklist := make(chan *node)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- &node{0, flag.Args()} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		node := <-worklist
		for _, link := range node.links {
			if !seen[link] {
				fmt.Println(link)
				seen[link] = true
				n++
				go func(depth int, link string) {
					worklist <- crawl(depth, link)
				}(node.depth, link)
			}
		}
	}
}
