package main

import "fmt"

// Dynamic Programming 1
// f(n, k) = f(n-1, k) + for each of f(n-1,k-1) combine with n
func combine(n int, k int) [][]int {
	res := [][]int{}
	if k == 1 {
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		return res
	}

	if k == n {
		tmp := []int{}
		for i := 1; i <= n; i++ {
			tmp = append(tmp, i)
		}
		res = append(res, tmp)
		return res
	}

	res = combine(n-1, k)
	left := combine(n-1, k-1)
	for _, v := range left {
		res = append(res, append(v, n))
	}
	return res
}

// Dynamic Programming 2, LeetCode Memory Limit Exceed
func combine2(n int, k int) [][]int {
	res := make([][][]int, k+1)
	res[0] = [][]int{}
	for i := 1; i <= n; i++ {
		res[1] = append(res[1], []int{i})
	}

	for i := 2; i <= k; i++ {
		fmt.Println(res[i])
		for _, v := range res[i-1] {
			last := v[len(v)-1]
			if last < n {
				for j := last + 1; j <= n; j++ {
					// note: must do a copy of the value
					pre := append([]int{}, v...)
					res[i] = append(res[i], append(pre, j))
				}
			}
		}
	}

	return res[k]
}
