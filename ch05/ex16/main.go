package main

import "strings"

func Join(sep string, a ...string) string {
	return strings.Join(a, sep)
}
