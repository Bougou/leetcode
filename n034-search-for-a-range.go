package main

func searchRange(nums []int, target int) []int {
	i := binarySearch(nums, target, 0, len(nums)-1)
	if i == -1 {
		return []int{-1, -1}
	}

	var l, r int
	for l = i; l >= 0 && nums[l] == target; l-- {
	}
	for r = i; r < len(nums) && nums[r] == target; r++ {
	}

	return []int{l + 1, r - 1}
}

func binarySearch(nums []int, target int, left int, right int) int {
	if len(nums) == 0 || left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] > target {
		return binarySearch(nums, target, left, mid-1)
	} else if nums[mid] < target {
		return binarySearch(nums, target, mid+1, right)
	} else {
		return mid
	}
}
