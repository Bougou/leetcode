package main

func twoSum(nums []int, target int) []int {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+nums[i] == target {
				return []int{j, i}
			}
		}
	}

	return []int{}
}
