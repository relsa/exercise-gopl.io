package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDiffBit(ha, hb [32]byte) int {
	var sum int
	var diff [32]byte

	for i := 0; i < 32; i++ {
		diff[i] = ha[i] ^ hb[i]
	}

	for _, b := range diff {
		sum += int(pc[b])
	}

	return sum
}
