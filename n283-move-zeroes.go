package main

func moveZeroes(nums []int) {
	// numNonZero
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// swap nums[j] and nums[i]

			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp

			j++
		}
	}
}
