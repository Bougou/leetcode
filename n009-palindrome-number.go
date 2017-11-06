package main

import (
	"math"
)
func isPalindrome(x int) bool {
	if x < 0 {
			return false
	}
	
	numLen := 1
	for y := x; y / 10 > 0; y = y / 10 {
			numLen++
	}
	
	if numLen == 1 {
			return true
	}
	
	for i := numLen-1 ; i > 0 ; i -= 2 {
			if x / int(math.Pow10(i)) !=  x % 10 {
					return false
			} else {
					x = (x % int(math.Pow10(i))) / 10
			}
	}
	
	return true
}