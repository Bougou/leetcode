package main

import (
	"fmt"
)

// Dynamic Programming
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
	}

	// the first element
	if obstacleGrid[0][0] == 1 {
		path[0][0] = 0
	} else {
		path[0][0] = 1
	}
	fmt.Println(path)

	// the top line
	for i := 1; i < n; i++ {
		if path[0][i-1] == 0 {
			path[0][i] = 0
		} else {
			if obstacleGrid[0][i] == 1 {
				path[0][i] = 0
			} else {
				path[0][i] = 1
			}
		}
	}
	fmt.Println(path)

	// the left line
	for i := 1; i < m; i++ {
		if path[i-1][0] == 0 {
			path[i][0] = 0
		} else {
			if obstacleGrid[i][0] == 1 {
				path[i][0] = 0
			} else {
				path[i][0] = 1
			}
		}
	}
	fmt.Println(path)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				path[i][j] = 0
			} else {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}
		}
	}

	fmt.Println(path)

	return path[m-1][n-1]
}
