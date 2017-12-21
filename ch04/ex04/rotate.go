package main

func rotate(s []int, n int) {
	l := len(s)
	g := gcd(l, n)

	for i := 0; i < g; i++ {
		for j := (i + n) % l; j != i; j = (j + n) % l {
			s[i], s[j] = s[j], s[i]
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
