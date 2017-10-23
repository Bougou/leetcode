package main

// mySqrt(9) -> 3
// mySqrt(4) -> 2
// mySqrt(3) -> 1
// mySqrt(2) -> 1
// mySqrt(1) -> 1
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	low := 1
	high := x
	for low <= high {
		mid := (low + high) / 2
		// cannot use: mid * mid > x , this might overflow
		if mid > x/mid {
			high = mid - 1
		} else if mid < x/mid {
			low = mid + 1
		} else {
			return mid
		}
	}

	return high
}
