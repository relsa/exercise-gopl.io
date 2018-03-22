package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByShifting(x uint64) int {
	var cnt int
	for i := uint(0); i < 64; i++ {
		cnt += int((x >> i) & 1)
	}
	return cnt
}

func PopCountByClearing(x uint64) int {
	var cnt int
	for x > 0 {
		x &= x - 1
		cnt++
	}
	return cnt
}
