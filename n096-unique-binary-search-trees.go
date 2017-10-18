package main

// https://discuss.leetcode.com/topic/8398/dp-solution-in-6-lines-with-explanation-f-i-n-g-i-1-g-n-i/5

// Dynamic Programming
func numTrees(n int) int {
	nums := make([]int, n+1)

	nums[0] = 1 // for empty tree
	nums[1] = 1 // for only one root

	for i := 2; i <= n; i++ {
		nums[i] = 0
		// nums[n] = nums[0]*nums[n-1] + nums[1]*nums[n-2]+ ... + nums[n-1]*nums[0]
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}

	return nums[n]
}
