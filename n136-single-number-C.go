package main

// This is classical interview question
// As we know, the same number XOR together will be 0,
// So, XOR all of numbers, the result is the number which only appears once. 

func singleNumber(nums []int) int {
	s := nums[0]

	for i := 1; i < len(nums); i++ {
		s ^= nums[i]
	}

	return s
}
