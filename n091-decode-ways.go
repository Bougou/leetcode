// A message containing letters from A-Z is being encoded to numbers using the following mapping:

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given an encoded message containing digits, determine the total number of ways to decode it.

// For example,
// Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).

// The number of ways decoding "12" is 2.

// Dynamic Programming
package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	fmt.Println(s)
	nums := make([]int, len(s)+1)
	nums[0] = 1
	nums[1] = 1

	// i mean the index of s, the Decoding number is stored in nums[i+1]
	for i := 1; i < len(s); i++ {
		nums[i+1] = nums[i]

		d, _ := strconv.Atoi(s[i-1 : i+1])
		fmt.Println(d)
		if d <= 26 {
			nums[i+1] += nums[i-1]
		}
		fmt.Println(nums)
	}

	return nums[len(s)]
}
