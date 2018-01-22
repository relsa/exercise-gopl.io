package main

import (
	"regexp"
)

func expand(s string, f func(string) string) string {
	reg := regexp.MustCompile(`\$\w*`)
	return reg.ReplaceAllStringFunc(s, func(x string) string {
		return f(x[1:])
	})
}
