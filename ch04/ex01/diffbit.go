package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDiffBit(ha, hb [32]byte) int {
	var sum int

	for i := 0; i < 32; i++ {
		sum += int(pc[ha[i]^hb[i]])
	}

	return sum
}
