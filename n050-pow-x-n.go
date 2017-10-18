package main

func myPow(x float64, n int) float64 {
	if n < 0 {
			n = -n
			x = 1/x
	}
	
	var res float64 = 1.0
	for n > 0 {
			// n is odd
			if n & 1 == 1 {
					res *= x
			}
			x *= x
			n >>= 1
	}
	
	return res

}