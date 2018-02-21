// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestString(t *testing.T) {
	ts := []struct {
		data []int
		want string
	}{
		{[]int{}, ""},
		{[]int{1, 2, 3, 4, 5}, "1 2 3 4 5"},
	}

	for _, tc := range ts {
		var tr *tree
		for _, v := range tc.data {
			tr = add(tr, v)
		}

		if got := tr.String(); got != tc.want {
			t.Errorf("got %q, want %q", got, tc.want)
		}
	}
}
