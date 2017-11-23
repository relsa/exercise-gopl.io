package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var cnt int

	for i := 0; i < 64; i++ {
		cnt += int((x >> uint64(i)) & uint64(1))
	}

	return cnt
}
