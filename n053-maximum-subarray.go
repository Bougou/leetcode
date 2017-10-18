package main

// O(n^2)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	//begin := 0
	//end := 0

	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i; j >= 0; j-- {
			if j != i {
				sum += nums[j]
			}

			if sum > maxSum {
				maxSum = sum
				//end = i
				//begin = j
			}
		}
	}

	return maxSum
}

// O(n)
func maxSubArray2(nums []int) int {
	maxSum := nums[0]

	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if preSum > maxSum {
			maxSum = preSum
		}

		//如果一段序列的和值小于0，这段序列的和值就没有必要加给后面的元素。
		if preSum < 0 {
			preSum = 0
		}
	}

	return maxSum
}
