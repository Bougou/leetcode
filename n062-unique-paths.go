package main

import "fmt"

// Dynamic Programming
func uniquePaths(m int, n int) int {
	// every element represent the unique paths to reach to this location.
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
	}

	path[0][0] = 1

	// the top line
	for i := 1; i < n; i++ {
		path[0][i] = 1
	}

	// the left line
	for i := 1; i < m; i++ {
		path[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i][j-1] + path[i-1][j]
		}
	}

	fmt.Println(path)
	return path[m-1][n-1]
}
