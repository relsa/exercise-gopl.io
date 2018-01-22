package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"}, // loop
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, ok := topoSort(prereqs)
	if !ok {
		log.Fatal("the graph has loops")
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) (order []string, ok bool) {
	seen := make(map[string]bool)
	processing := make(map[string]bool)

	var visitAll func(items []string)
	ok = true

	visitAll = func(items []string) {
		for _, item := range items {
			if processing[item] {
				ok = false
				return
			}
			if !seen[item] {
				processing[item] = true
				visitAll(m[item])
				order = append(order, item)

				processing[item] = false
				seen[item] = true
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order, ok
}
