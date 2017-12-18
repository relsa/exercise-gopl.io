package main

import (
	"strings"
)

func anagram(sa, sb string) bool {
	for _, ra := range sa {
		b := strings.ContainsRune(sb, ra)
		if !b {
			return false
		}

		sb = strings.Replace(sb, string(ra), "", 1)
	}

	return sb == ""
}
