package main

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		vals    []int
		isValid bool
		want    int
	}{
		{nil, false, math.MinInt32},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3}, true, 3},
	}

	for _, test := range tests {
		got, err := max(test.vals...)
		if err == nil {
			if !test.isValid {
				t.Errorf("case %v: invalid args but no errors", test.vals)
			}
			if got != test.want {
				t.Errorf("case %v: got %d, want %d", test.vals, got, test.want)
			}
		} else {
			if test.isValid {
				t.Errorf("case %v: valid args but errors raised", test.vals)
			}
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		vals    []int
		isValid bool
		want    int
	}{
		{nil, false, math.MaxInt32},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3}, true, 1},
	}

	for _, test := range tests {
		got, err := min(test.vals...)
		if err == nil {
			if !test.isValid {
				t.Errorf("case %v: invalid args but no errors", test.vals)
			}
			if got != test.want {
				t.Errorf("case %v: got %d, want %d", test.vals, got, test.want)
			}
		} else {
			if test.isValid {
				t.Errorf("case %v: valid args but errors raised", test.vals)
			}
		}
	}
}

func TestMax1(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 3},
	}

	for _, test := range tests {
		got := max1(test.vals[0], test.vals[1:]...)
		if got != test.want {
			t.Errorf("case %v: got %d, want %d", test.vals, got, test.want)
		}
	}
}

func TestMin1(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 1},
	}

	for _, test := range tests {
		got := min1(test.vals[0], test.vals[1:]...)
		if got != test.want {
			t.Errorf("case %v: got %d, want %d", test.vals, got, test.want)
		}
	}
}
