// Follow up for "Remove Duplicates":
// What if duplicates are allowed at most twice?

// For example,
// Given sorted array nums = [1,1,1,2,2,3],

// Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.

package main

func removeDuplicates(nums []int) int {
	i := 0
	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}
	return i
}
