package main

/*
 *  sum + m*(n-1) = x*n
 * 	x = minNum + m
 *  n*Num + m = sum
 */

func minMoves(nums []int) int {
	minNum := nums[0]
	sum := 0

	for _, num := range nums {
		sum += num
		if num < minNum {
			minNum = num
		}
	}

	return sum - len(nums)*minNum
}
