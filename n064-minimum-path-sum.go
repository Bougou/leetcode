package main

import (
	"fmt"
)

// Dynamic Programming
func minPathSum(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])

	minSum := make([][]int, r)

	for i := range minSum {
		minSum[i] = make([]int, c)
	}

	// initialize the top left node.
	minSum[0][0] = grid[0][0]

	// the top most line
	for i := 1; i < c; i++ {
		minSum[0][i] = minSum[0][i-1] + grid[0][i]
	}

	// the left most line
	for j := 1; j < r; j++ {
		minSum[j][0] = minSum[j-1][0] + grid[j][0]
	}
	fmt.Println(minSum)

	for m := 1; m < r; m++ {
		for n := 1; n < c; n++ {
			// for each node in the matrix, the path is from its top or its left.
			min := minSum[m-1][n]
			if minSum[m][n-1] < minSum[m-1][n] {
				min = minSum[m][n-1]
			}
			minSum[m][n] = min + grid[m][n]
		}
	}
	fmt.Println(minSum)

	return minSum[r-1][c-1]
}
