package main

// recursive way
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

//
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	pre1 := 1
	pre2 := 2
	var res int
	for i := 3; i <= n; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}

	return res
}
