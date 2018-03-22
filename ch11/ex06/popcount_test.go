// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import (
	"testing"
)

func benchmarkPopCount(b *testing.B, v uint64) {
	for i := 0; i < b.N; i++ {
		PopCount(v)
	}
}

func benchmarkPopCountByShifting(b *testing.B, v uint64) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(v)
	}
}

func benchmarkPopCountByClearing(b *testing.B, v uint64) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(v)
	}
}

func BenchmarkPopCount_0(b *testing.B)           { benchmarkPopCount(b, 0x0) }
func BenchmarkPopCountByShifting_0(b *testing.B) { benchmarkPopCountByShifting(b, 0x0) }
func BenchmarkPopCountByClearing_0(b *testing.B) { benchmarkPopCountByClearing(b, 0x0) }

func BenchmarkPopCount_F(b *testing.B)           { benchmarkPopCount(b, 0xFFFFFFFF) }
func BenchmarkPopCountByShifting_F(b *testing.B) { benchmarkPopCountByShifting(b, 0xFFFFFFFF) }
func BenchmarkPopCountByClearing_F(b *testing.B) { benchmarkPopCountByClearing(b, 0xFFFFFFFF) }

func BenchmarkPopCount_1F(b *testing.B)           { benchmarkPopCount(b, 0x1234567890ABCDEF) }
func BenchmarkPopCountByShifting_1F(b *testing.B) { benchmarkPopCountByShifting(b, 0x1234567890ABCDEF) }
func BenchmarkPopCountByClearing_1F(b *testing.B) { benchmarkPopCountByClearing(b, 0x1234567890ABCDEF) }
