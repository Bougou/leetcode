package main

import "math"

func titleToNumber(s string) int {
	var sum int

	for i, c := range []rune(s) {
		k := int(c - 'A' + 1)
		exp := float64(len(s) - i - 1)
		sum += k * int(math.Pow(26, exp))
	}

	return sum
}
