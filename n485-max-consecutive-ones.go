package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cons := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cons++
			if cons > max {
				max = cons
			}
		} else {
			cons = 0
		}
	}

	return max
}
