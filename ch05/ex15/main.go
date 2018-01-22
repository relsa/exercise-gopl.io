package main

import "fmt"
import "math"

func max(vals ...int) (int, error) {
	m := math.MinInt32

	if len(vals) == 0 {
		return m, fmt.Errorf("no args")
	}

	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m, nil
}

func min(vals ...int) (int, error) {
	m := math.MaxInt32

	if len(vals) == 0 {
		return m, fmt.Errorf("no args")
	}

	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m, nil
}

func max1(v int, vals ...int) int {
	m := v
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m
}

func min1(v int, vals ...int) int {
	m := v
	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m
}
