// Given a set of distinct integers, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

// For example,
// If nums = [1,2,3], a solution is:

// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

// Recursive Solution
// Suppose S(n) represents the subsets of [1,2,3,...,n]
// The idea here is S(n) = S(n-1) + [for s in S(n-1), append(s, n)]
package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	last := nums[len(nums)-1]
	// recursive call
	subset := subsets(nums[:len(nums)-1])

	res := [][]int{}
	res = subset
	for _, s := range subset {
		// must do a copy
		tmp := append([]int{}, s...)
		res = append(res, append(tmp, last))
	}
	fmt.Println(res)
	return res
}
