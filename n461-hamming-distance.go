package main

import "fmt"

// 计算出两个数的异或值的二进制中包含的1的数量
func hammingDistance(x int, y int) int {
	fmt.Println("#461-hamming-distance")
	xor := x ^ y
	fmt.Printf("%b\n%b\n%b\n", x, y, xor)

	count := 0

	var i uint
	// an int value is 32 bit
	for i = 0; i < 32; i++ {
		count += (xor >> i) & 1
	}

	return count
}
