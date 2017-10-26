// Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

// For example,
// If nums = [1,2,2], a solution is:

// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]

// see: https://leetcode.com/problems/subsets-ii/discuss/
package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.IntSlice(nums).Sort()

	res := [][]int{
		[]int{},
	}
	for i := 0; i < len(nums); {
		count := 0
		for i+count < len(nums) && nums[i+count] == nums[i] {
			count++
		}

		prevLen := len(res)
		for j := 0; j < prevLen; j++ {
			for c := 1; c <= count; c++ {
				tmp := append([]int{}, res[j]...)
				for k := 0; k < c; k++ {
					tmp = append(tmp, nums[i])
				}
				res = append(res, tmp)
			}
		}
		i += count
	}

	return res
}
