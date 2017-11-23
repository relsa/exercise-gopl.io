package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var cnt int

	for x > 0 {
		x &= (x - uint64(1))
		cnt++
	}

	return cnt
}
